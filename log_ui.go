package main

import (
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

func DisplayLogUIHeadings(app *cview.Application) {
	// Create the logging box
	LogTable.SetBorders(true)
	timeHeading := cview.NewTableCell("Time")
	timeHeading.SetTextColor(tcell.ColorYellow.TrueColor())
	timeHeading.SetAlign(cview.AlignCenter)
	timeHeading.SetExpansion(1)
	LogTable.SetCell(0, 0, timeHeading)
	callHeading := cview.NewTableCell("Call")
	callHeading.SetTextColor(tcell.ColorYellow.TrueColor())
	callHeading.SetAlign(cview.AlignCenter)
	callHeading.SetExpansion(1)
	LogTable.SetCell(0, 1, callHeading)
	bandHeading := cview.NewTableCell("Band")
	bandHeading.SetTextColor(tcell.ColorYellow.TrueColor())
	bandHeading.SetAlign(cview.AlignCenter)
	bandHeading.SetExpansion(1)
	LogTable.SetCell(0, 2, bandHeading)
	modeHeading := cview.NewTableCell("Mode")
	modeHeading.SetTextColor(tcell.ColorYellow.TrueColor())
	modeHeading.SetAlign(cview.AlignCenter)
	modeHeading.SetExpansion(1)
	LogTable.SetCell(0, 3, modeHeading)
	parkHeading := cview.NewTableCell("Park")
	parkHeading.SetTextColor(tcell.ColorYellow.TrueColor())
	parkHeading.SetAlign(cview.AlignCenter)
	parkHeading.SetExpansion(1)
	LogTable.SetCell(0, 4, parkHeading)
}

func DisplayLogUIContacts(app *cview.Application) {
	// Do we have any contacts to display?
	if Op.NumContacts == 0 {
		return
	}
	// Compose our database query
	dbQuery := "SELECT time_on, call, band, mode, sig_info from log order by time_on desc limit 10;"
	result, err := Op.Database.Query(dbQuery)
	if err != nil {
		panic(err)
	}
	// Initialize our row counter
	row := 1
	// Slice to hold an individual row
	var rowSlice []string = make([]string, 5)
	// Iterate through the contacts
	for result.Next() {
		if err := result.Scan(&rowSlice[0], &rowSlice[1], &rowSlice[2], &rowSlice[3], &rowSlice[4]); err != nil {
			panic(err)
		}
		time := cview.NewTableCell(rowSlice[0])
		time.SetAlign(cview.AlignCenter)
		LogTable.SetCell(row, 0, time)
		call := cview.NewTableCell(rowSlice[1])
		call.SetAlign(cview.AlignCenter)
		LogTable.SetCell(row, 1, call)
		band := cview.NewTableCell(rowSlice[2])
		band.SetAlign(cview.AlignCenter)
		LogTable.SetCell(row, 2, band)
		mode := cview.NewTableCell(rowSlice[3])
		mode.SetAlign(cview.AlignCenter)
		LogTable.SetCell(row, 3, mode)
		park := cview.NewTableCell(rowSlice[4])
		park.SetAlign(cview.AlignCenter)
		LogTable.SetCell(row, 4, park)
		// Increment our row counter
		row++
	}
}
