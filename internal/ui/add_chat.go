// Copyright 2020 The Mellium Contributors.
// Use of this source code is governed by the BSD 2-clause
// license that can be found in the LICENSE file.
package ui

import (
	"github.com/rivo/tview"

	"mellium.im/xmpp/jid"
)

// addRoster creates a modal that asks for a JID to add to the roster.
func addRoster(addButton string, autocomplete []jid.JID, f func(jid.JID, string)) *Modal {
	mod := NewModal()
	mod.SetText("Add Contact")

	var inputJID jid.JID
	jidInput := jidInput(&inputJID, autocomplete)
	jidInput.SetLabel("Address")

	modForm := mod.Form()
	modForm.AddFormItem(jidInput)
	mod.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor).
		AddButtons([]string{cancelButton, addButton}).
		SetDoneFunc(func(_ int, buttonLabel string) {
			f(inputJID.Bare(), buttonLabel)
		})
	return mod
}