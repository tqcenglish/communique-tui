// Copyright 2018 The Mellium Contributors.
// Use of this source code is governed by the BSD 2-clause
// license that can be found in the LICENSE file.

package client

import (
	"context"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"mellium.im/communiqué/internal/logwriter"
	"mellium.im/communiqué/internal/ui"
	"mellium.im/sasl"
	"mellium.im/xmlstream"
	"mellium.im/xmpp"
	"mellium.im/xmpp/dial"
	"mellium.im/xmpp/jid"
	"mellium.im/xmpp/roster"
	"mellium.im/xmpp/stanza"
)

// New creates a new XMPP client but does not attempt to negotiate a session or
// send an initial presence, etc.
func New(timeout time.Duration, configPath, addr, keylogFile string, pane *ui.UI, xmlIn, xmlOut, logger, debug *log.Logger, getPass func(context.Context) (string, error)) *Client {
	var j jid.JID
	var err error
	if addr == "" {
		logger.Printf(`No user address specified, edit %q and add:

	jid="me@example.com"

`, configPath)
	} else {
		logger.Printf("User address: %q", addr)
		j, err = jid.Parse(addr)
		if err != nil {
			logger.Printf("Error parsing user address: %q", err)
		}
	}

	var keylog io.Writer
	if keylogFile != "" {
		keylog, err = os.OpenFile(keylogFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0400)
		if err != nil {
			logger.Printf("Error creating keylog file: %q", err)
		}
	}
	dialer := &dial.Dialer{
		TLSConfig: &tls.Config{
			ServerName:   j.Domain().String(),
			KeyLogWriter: keylog,
		},
	}

	c := &Client{
		timeout: timeout,
		addr:    j,
		dialer:  dialer,
		logger:  logger,
		debug:   debug,
		pane:    pane,
		getPass: getPass,
	}
	if xmlIn != nil {
		c.win = logwriter.New(xmlIn)
	}
	if xmlOut != nil {
		c.wout = logwriter.New(xmlOut)
	}

	pane.Offline()
	return c
}

func (c *Client) reconnect(ctx context.Context) error {
	if c.online {
		return nil
	}

	pass, err := c.getPass(ctx)
	if err != nil {
		return err
	}

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, c.timeout)
	defer cancel()

	conn, err := c.dialer.Dial(ctx, "tcp", c.addr)
	if err != nil {
		return fmt.Errorf("error dialing connection: %w", err)
	}

	negotiator := xmpp.NewNegotiator(xmpp.StreamConfig{
		Features: []xmpp.StreamFeature{
			xmpp.StartTLS(true, c.dialer.TLSConfig),
			xmpp.SASL("", pass, sasl.ScramSha256Plus, sasl.ScramSha1Plus, sasl.ScramSha256, sasl.ScramSha1),
			xmpp.BindResource(),
		},
		TeeIn:  c.win,
		TeeOut: c.wout,
	})
	c.Session, err = xmpp.NegotiateSession(ctx, c.addr.Domain(), c.addr, conn, false, negotiator)
	if err != nil {
		return fmt.Errorf("error negotiating session: %w", err)
	}

	c.online = true
	go func() {
		err := c.Serve(newXMPPHandler(c))
		if err != nil {
			c.logger.Printf("Error while handling XMPP streams: %q", err)
		}
		c.online = false
		c.pane.Offline()
		if err = conn.Close(); err != nil {
			c.logger.Printf("Error closing the connection: %q", err)
		}
	}()

	// TODO: should this be synchronous so that when we call reconnect we fail if
	// the roster isn't fetched?
	go func() {
		rosterCtx, rosterCancel := context.WithTimeout(context.Background(), c.timeout)
		defer rosterCancel()
		err = c.Roster(rosterCtx)
		if err != nil {
			c.logger.Printf("Error fetching roster: %q", err)
		}
	}()
	return nil
}

// Client represents an XMPP client.
type Client struct {
	*xmpp.Session
	timeout time.Duration
	pane    *ui.UI
	logger  *log.Logger
	debug   *log.Logger
	addr    jid.JID
	win     io.Writer
	wout    io.Writer
	dialer  *dial.Dialer
	getPass func(context.Context) (string, error)
	online  bool
}

// Online sets the status to online.
// The provided context is used if the client was previously offline and we
// have to re-establish the session, so if it includes a timeout make sure to
// account for the fact that we might reconnect.
func (c *Client) Online(ctx context.Context) {
	err := c.reconnect(ctx)
	if err != nil {
		c.logger.Println(err)
		return
	}

	err = c.Send(ctx, stanza.WrapPresence(jid.JID{}, stanza.AvailablePresence, nil))
	if err != nil {
		c.logger.Printf("Error sending online presence: %q", err)
		return
	}
	c.pane.Online()
}

// Roster requests the users contact list.
func (c *Client) Roster(ctx context.Context) error {
	iter := roster.Fetch(ctx, c.Session)
	defer func() {
		e := iter.Close()
		if e != nil {
			c.debug.Printf("Error closing roster stream: %q", e)
		}
	}()
	for iter.Next() {
		item := iter.Item()
		if item.Name == "" {
			item.Name = item.JID.Localpart()
		}
		if item.Name == "" {
			item.Name = item.JID.Domainpart()
		}
		c.pane.UpdateRoster(ui.RosterItem{Item: item})
	}
	err := iter.Err()
	if err != io.EOF {
		return err
	}

	return nil
}

// Away sets the status to away.
func (c *Client) Away(ctx context.Context) {
	err := c.reconnect(ctx)
	if err != nil {
		c.logger.Println(err)
		return
	}

	err = c.Send(
		ctx,
		stanza.WrapPresence(
			jid.JID{},
			stanza.AvailablePresence,
			xmlstream.Wrap(
				xmlstream.ReaderFunc(func() (xml.Token, error) {
					return xml.CharData("away"), io.EOF
				}),
				xml.StartElement{Name: xml.Name{Local: "show"}},
			)))
	if err != nil {
		c.logger.Printf("Error sending away presence: %q", err)
		return
	}
	c.pane.Away()
}

// Busy sets the status to busy.
func (c *Client) Busy(ctx context.Context) {
	err := c.reconnect(ctx)
	if err != nil {
		c.logger.Println(err)
		return
	}

	err = c.Send(
		ctx,
		stanza.WrapPresence(
			jid.JID{},
			stanza.AvailablePresence,
			xmlstream.Wrap(
				xmlstream.ReaderFunc(func() (xml.Token, error) {
					return xml.CharData("dnd"), io.EOF
				}),
				xml.StartElement{Name: xml.Name{Local: "show"}},
			)))
	if err != nil {
		c.logger.Printf("Error sending busy presence: %q", err)
		return
	}
	c.pane.Busy()
}

// Offline logs the client off.
func (c *Client) Offline() {
	defer c.pane.Offline()
	if !c.online {
		return
	}

	err := c.SetCloseDeadline(time.Now().Add(30 * time.Second))
	if err != nil {
		c.debug.Printf("Error setting close deadline: %q", err)
		// Don't return; we still want to attempt to close the connection.
	}
	err = c.Close()
	if err != nil {
		c.logger.Printf("Error logging off: %q", err)
	}
	c.online = false
}