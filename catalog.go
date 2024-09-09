// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en_US": &dictionary{index: en_USIndex, data: en_USData},
	}
	fallback := language.MustParse("en-US")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"%v\n\nTry running '%s -config' to generate a default config file.": 33,
	"Add":                              118,
	"Add Contact":                      99,
	"Address":                          100,
	"Are you sure you want to quit?":   114,
	"Away":                             112,
	"Busy":                             113,
	"Cancel":                           21,
	"Channels":                         102,
	"Commands":                         119,
	"Complete":                         20,
	"Conversation":                     103,
	"Conversations":                    105,
	"DEBUG":                            25,
	"Error closing roster stream: %q":  97,
	"Error closing the connection: %q": 86,
	"Error encoding default config as TOML: %v": 32,
	"Error going offline: %q":                   85,
	"Error while handling XMPP streams: %q":     84,
	"Exec":                                      121,
	"Join":                                      116,
	"Join Channel":                              117,
	"Loading commands…":                         120,
	"Logs":                                      106,
	"Name":                                      101,
	"Next":                                      19,
	"Offline":                                   110,
	"Online":                                    111,
	"Prev":                                      18,
	"Quit":                                      115,
	"RECV":                                      26,
	"Remove":                                    107,
	"Remove this channel?":                      109,
	"Remove this contact from your roster?":     108,
	"SENT":                                      27,
	"To fix this, contact your server administrator and ask them to enable %q": 94,
	"Usage of communiqué:\n\n": 24,
	"Your server does not support bookmark unification, an important feature that stops newer clients from seeing a different list of chat rooms than older clients that do not yet support the latest features.": 95,
	"account %q not found in config file":                                         36,
	"caps cache hit for %s: %s:%s":                                                15,
	"caps cache miss for %s: %s:%s, %[2]s:%[4]s":                                  16,
	"could not get the upload services: %v":                                       73,
	"could not upload %q: %v":                                                     74,
	"error adding roster item %s: %v":                                             61,
	"error bootstraping history for %s: %v":                                       3,
	"error closing bookmarks stream: %v":                                          96,
	"error closing commands iter for %q: %v":                                      56,
	"error closing config file: %v":                                               35,
	"error copying early log data to output buffer: %q":                           40,
	"error creating keylog file: %q":                                              47,
	"error discovering bookmarks support: %v":                                     63,
	"error enabling carbons: %q":                                                  90,
	"error executing command %q on %q: %v":                                        53,
	"error fetching bookmarks: %q":                                                92,
	"error fetching commands for %q: %v":                                          55,
	"error fetching earliest message info for %v from database: %v":               79,
	"error fetching history after %s for %s: %v":                                  2,
	"error fetching info from cache: %v":                                          13,
	"error fetching roster: %q":                                                   91,
	"error fetching scrollback for %v: %v":                                        82,
	"error fetching version information: %v":                                      93,
	"error going offline: %v":                                                     60,
	"error inserting entity capbailities hash: %v":                                11,
	"error iterating over roster items: %v":                                       4,
	"error joining room %s: %v":                                                   78,
	"error loading chat: %v":                                                      75,
	"error logging to pane: %v":                                                   39,
	"error marking message %q as received: %v":                                    6,
	"error occured during service discovery: %v":                                  87,
	"error opening database: %v":                                                  38,
	"error parsing config file: %v":                                               34,
	"error parsing main account as XMPP address: %v":                              37,
	"error parsing timeout, defaulting to 30s: %q":                                46,
	"error parsing user address: %q":                                              45,
	"error publishing bookmark %s: %v":                                            65,
	"error publishing legacy bookmark %s: %v":                                     64,
	"error querying database for last seen messages: %v":                          1,
	"error querying history for %s: %v":                                           23,
	"error removing bookmark %s: %v":                                              67,
	"error removing legacy bookmark %s: %v":                                       66,
	"error removing roster item %s: %v":                                           62,
	"error retrieving roster version, falling back to full roster fetch: %v":      48,
	"error running password command, falling back to prompt: %v":                  42,
	"error saving entity caps to the database: %v":                                17,
	"error saving sent message to history: %v":                                    72,
	"error scrollback for %v: %v":                                                 83,
	"error sending message: %v":                                                   71,
	"error sending presence pre-approval to %s: %v":                               68,
	"error sending presence request to %s: %v":                                    69,
	"error setting away status: %v":                                               57,
	"error setting busy status: %v":                                               59,
	"error setting online status: %v":                                             58,
	"error showing next command for %q: %v":                                       54,
	"error updating roster version: %v":                                           5,
	"error updating to roster ver %q: %v":                                         0,
	"error when closing the items iterator: %v":                                   88,
	"error while picking files: %v":                                               104,
	"error writing history message to chat: %v":                                   9,
	"error writing history to database: %v":                                       10,
	"error writing history: %v":                                                   22,
	"error writing message to database: %v":                                       8,
	"error writing received message to chat: %v":                                  7,
	"executing command: %+v":                                                      52,
	"failed to read stderr of the notification subprocess: %v":                    122,
	"failed to run notification command: %v":                                      123,
	"falling back to network query…":                                              14,
	"feature discovery failed for %q: %v":                                         89,
	"fetching scrollback before %v for %v…":                                       81,
	"got signal: %v":                                                              51,
	"initial login failed: %v":                                                    49,
	"invalid nick %s in config: %v":                                               76,
	"joining room %v…":                                                            77,
	"logged in as: %q":                                                            50,
	"no scrollback for %v":                                                        80,
	"no user address specified, edit %q and add:\n\n\tjid=\"me@example.com\"\n\n": 43,
	"notification subprocess failed: %v\n%s":                                      124,
	"override the account set in the config file":                                 29,
	"possibly spoofed history message from %s":                                    98,
	"print a default config file to stdout":                                       31,
	"print this help message":                                                     30,
	"running command: %q":                                                         41,
	"the config file to load":                                                     28,
	"unrecognized client event: %T(%[1]q)":                                        12,
	"unrecognized ui event: %T(%[1]q)":                                            70,
	"user address: %q":                                                            44,
}

var en_USIndex = []uint32{ // 126 elements
	// Entry 0 - 1F
	0x00000000, 0x0000002a, 0x00000060, 0x00000094,
	0x000000c0, 0x000000e9, 0x0000010e, 0x0000013d,
	0x0000016b, 0x00000194, 0x000001c1, 0x000001ea,
	0x0000021a, 0x00000242, 0x00000268, 0x00000289,
	0x000002af, 0x000002e3, 0x00000313, 0x00000318,
	0x0000031d, 0x00000326, 0x0000032d, 0x0000034a,
	0x00000372, 0x0000038e, 0x00000394, 0x00000399,
	0x0000039e, 0x000003b6, 0x000003e2, 0x000003fa,
	// Entry 20 - 3F
	0x00000420, 0x0000044d, 0x00000493, 0x000004b4,
	0x000004d5, 0x000004fc, 0x0000052e, 0x0000054c,
	0x00000569, 0x0000059e, 0x000005b5, 0x000005f3,
	0x0000063f, 0x00000653, 0x00000675, 0x000006a5,
	0x000006c7, 0x00000711, 0x0000072d, 0x00000741,
	0x00000753, 0x0000076d, 0x0000079b, 0x000007c7,
	0x000007f0, 0x0000081d, 0x0000083e, 0x00000861,
	0x00000882, 0x0000089d, 0x000008c3, 0x000008eb,
	// Entry 40 - 5F
	0x00000916, 0x00000944, 0x0000096b, 0x00000997,
	0x000009bc, 0x000009f0, 0x00000a1f, 0x00000a43,
	0x00000a60, 0x00000a8c, 0x00000ab5, 0x00000ad3,
	0x00000aed, 0x00000b11, 0x00000b27, 0x00000b47,
	0x00000b8b, 0x00000ba3, 0x00000bd1, 0x00000bfc,
	0x00000c1e, 0x00000c47, 0x00000c62, 0x00000c86,
	0x00000cb4, 0x00000ce1, 0x00000d0b, 0x00000d29,
	0x00000d46, 0x00000d66, 0x00000d90, 0x00000ddc,
	// Entry 60 - 7F
	0x00000ea8, 0x00000ece, 0x00000ef1, 0x00000f1d,
	0x00000f29, 0x00000f31, 0x00000f36, 0x00000f3f,
	0x00000f4c, 0x00000f6d, 0x00000f7b, 0x00000f80,
	0x00000f87, 0x00000fad, 0x00000fc2, 0x00000fca,
	0x00000fd1, 0x00000fd6, 0x00000fdb, 0x00000ffa,
	0x00000fff, 0x00001004, 0x00001011, 0x00001015,
	0x0000101e, 0x00001032, 0x00001037, 0x00001073,
	0x0000109d, 0x000010c9,
} // Size: 528 bytes

const en_USData string = "" + // Size: 4297 bytes
	"\x02error updating to roster ver %[1]q: %[2]v\x02error querying database" +
	" for last seen messages: %[1]v\x02error fetching history after %[1]s for" +
	" %[2]s: %[3]v\x02error bootstraping history for %[1]s: %[2]v\x02error it" +
	"erating over roster items: %[1]v\x02error updating roster version: %[1]v" +
	"\x02error marking message %[1]q as received: %[2]v\x02error writing rece" +
	"ived message to chat: %[1]v\x02error writing message to database: %[1]v" +
	"\x02error writing history message to chat: %[1]v\x02error writing histor" +
	"y to database: %[1]v\x02error inserting entity capbailities hash: %[1]v" +
	"\x02unrecognized client event: %[1]T(%[1]q)\x02error fetching info from " +
	"cache: %[1]v\x02falling back to network query…\x02caps cache hit for %[1" +
	"]s: %[2]s:%[3]s\x02caps cache miss for %[1]s: %[2]s:%[3]s, %[2]s:%[4]s" +
	"\x02error saving entity caps to the database: %[1]v\x02Prev\x02Next\x02C" +
	"omplete\x02Cancel\x02error writing history: %[1]v\x02error querying hist" +
	"ory for %[1]s: %[2]v\x04\x00\x02\x0a\x0a\x16\x02Usage of communiqué:\x02" +
	"DEBUG\x02RECV\x02SENT\x02the config file to load\x02override the account" +
	" set in the config file\x02print this help message\x02print a default co" +
	"nfig file to stdout\x02Error encoding default config as TOML: %[1]v\x02%" +
	"[1]v\x0a\x0aTry running '%[2]s -config' to generate a default config fil" +
	"e.\x02error parsing config file: %[1]v\x02error closing config file: %[1" +
	"]v\x02account %[1]q not found in config file\x02error parsing main accou" +
	"nt as XMPP address: %[1]v\x02error opening database: %[1]v\x02error logg" +
	"ing to pane: %[1]v\x02error copying early log data to output buffer: %[1" +
	"]q\x02running command: %[1]q\x02error running password command, falling " +
	"back to prompt: %[1]v\x04\x00\x02\x0a\x0aF\x02no user address specified," +
	" edit %[1]q and add:\x0a\x0a\x09jid=\x22me@example.com\x22\x02user addre" +
	"ss: %[1]q\x02error parsing user address: %[1]q\x02error parsing timeout," +
	" defaulting to 30s: %[1]q\x02error creating keylog file: %[1]q\x02error " +
	"retrieving roster version, falling back to full roster fetch: %[1]v\x02i" +
	"nitial login failed: %[1]v\x02logged in as: %[1]q\x02got signal: %[1]v" +
	"\x02executing command: %+[1]v\x02error executing command %[1]q on %[2]q:" +
	" %[3]v\x02error showing next command for %[1]q: %[2]v\x02error fetching " +
	"commands for %[1]q: %[2]v\x02error closing commands iter for %[1]q: %[2]" +
	"v\x02error setting away status: %[1]v\x02error setting online status: %[" +
	"1]v\x02error setting busy status: %[1]v\x02error going offline: %[1]v" +
	"\x02error adding roster item %[1]s: %[2]v\x02error removing roster item " +
	"%[1]s: %[2]v\x02error discovering bookmarks support: %[1]v\x02error publ" +
	"ishing legacy bookmark %[1]s: %[2]v\x02error publishing bookmark %[1]s: " +
	"%[2]v\x02error removing legacy bookmark %[1]s: %[2]v\x02error removing b" +
	"ookmark %[1]s: %[2]v\x02error sending presence pre-approval to %[1]s: %[" +
	"2]v\x02error sending presence request to %[1]s: %[2]v\x02unrecognized ui" +
	" event: %[1]T(%[1]q)\x02error sending message: %[1]v\x02error saving sen" +
	"t message to history: %[1]v\x02could not get the upload services: %[1]v" +
	"\x02could not upload %[1]q: %[2]v\x02error loading chat: %[1]v\x02invali" +
	"d nick %[1]s in config: %[2]v\x02joining room %[1]v…\x02error joining ro" +
	"om %[1]s: %[2]v\x02error fetching earliest message info for %[1]v from d" +
	"atabase: %[2]v\x02no scrollback for %[1]v\x02fetching scrollback before " +
	"%[1]v for %[2]v…\x02error fetching scrollback for %[1]v: %[2]v\x02error " +
	"scrollback for %[1]v: %[2]v\x02Error while handling XMPP streams: %[1]q" +
	"\x02Error going offline: %[1]q\x02Error closing the connection: %[1]q" +
	"\x02error occured during service discovery: %[1]v\x02error when closing " +
	"the items iterator: %[1]v\x02feature discovery failed for %[1]q: %[2]v" +
	"\x02error enabling carbons: %[1]q\x02error fetching roster: %[1]q\x02err" +
	"or fetching bookmarks: %[1]q\x02error fetching version information: %[1]" +
	"v\x02To fix this, contact your server administrator and ask them to enab" +
	"le %[1]q\x02Your server does not support bookmark unification, an import" +
	"ant feature that stops newer clients from seeing a different list of cha" +
	"t rooms than older clients that do not yet support the latest features." +
	"\x02error closing bookmarks stream: %[1]v\x02Error closing roster stream" +
	": %[1]q\x02possibly spoofed history message from %[1]s\x02Add Contact" +
	"\x02Address\x02Name\x02Channels\x02Conversation\x02error while picking f" +
	"iles: %[1]v\x02Conversations\x02Logs\x02Remove\x02Remove this contact fr" +
	"om your roster?\x02Remove this channel?\x02Offline\x02Online\x02Away\x02" +
	"Busy\x02Are you sure you want to quit?\x02Quit\x02Join\x02Join Channel" +
	"\x02Add\x02Commands\x02Loading commands…\x02Exec\x02failed to read stder" +
	"r of the notification subprocess: %[1]v\x02failed to run notification co" +
	"mmand: %[1]v\x02notification subprocess failed: %[1]v\x0a%[2]s"

	// Total table size 4825 bytes (4KiB); checksum: C85D7D17
