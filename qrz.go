package main

import (
	"context"
	"fmt"

	"code.rocketnine.space/tslocum/cview"
	"github.com/k0swe/qrz-api"
)

func QrzLookup(app *cview.Application) {
	ctx := context.Background()
	user := Op.QrzUsername
	pw := Op.QrzPassword
	lookup := WorkedCallsignInput.GetText()
	lookupResp, err := qrz.Lookup(ctx, &user, &pw, &lookup)
	if err != nil {
		text := StatusBox.GetText(false)
		text += fmt.Sprintf("%s\n", err)
		app.QueueUpdateDraw(func() {
			StatusBox.SetText(text)
		})
		return
	}
	if lookupResp.Callsign == (qrz.Callsign{}) {
		text := StatusBox.GetText(false)
		text += fmt.Sprintf("No results found for %s\n", lookup)
		app.QueueUpdateDraw(func() {
			StatusBox.SetText(text)
		})
		return
	} else {
		// Get our existing StatusBox text
		text := StatusBox.GetText(false)
		// Add our new text
		text += fmt.Sprintf("%s: %s %s, %s %s %s\n",
			lookupResp.Callsign.Call,
			lookupResp.Callsign.Fname,
			lookupResp.Callsign.Name,
			lookupResp.Callsign.Addr2,
			lookupResp.Callsign.State,
			lookupResp.Callsign.Country,
		)
		// Set our new text
		app.QueueUpdateDraw(func() {
			StatusBox.SetText(text)
		})
		return
	}
}
