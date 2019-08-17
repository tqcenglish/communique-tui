// Copyright 2018 The Mellium Contributors.
// Use of this source code is governed by the BSD 2-clause
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"log"

	"mellium.im/communiqué/internal/client"
	"mellium.im/communiqué/internal/ui"
)

func newUIHandler(c *client.Client, logger, debug *log.Logger) func(ui.Event) {
	return func(e ui.Event) {
		switch e {
		case ui.GoAway:
			go c.Away(context.TODO())
		case ui.GoOnline:
			go c.Online(context.TODO())
		case ui.GoBusy:
			go c.Busy(context.TODO())
		case ui.GoOffline:
			go c.Offline()
		default:
			debug.Printf("Unrecognized event: %q", e)
		}
	}
}
