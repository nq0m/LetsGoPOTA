package main

import (
	"strconv"

	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

// All our input fields for contacts made
var WorkedCallsignInput *cview.InputField = cview.NewInputField()
var SentReportInput *cview.InputField = cview.NewInputField()
var RcvdReportInput *cview.InputField = cview.NewInputField()
var WorkedParkInput *cview.InputField = cview.NewInputField()
var CommentsInput *cview.InputField = cview.NewInputField()
var FreqInput *cview.InputField = cview.NewInputField()
var BandDropDown *cview.DropDown = cview.NewDropDown()
var ModeDropDown *cview.DropDown = cview.NewDropDown()
var MyPowerInput *cview.InputField = cview.NewInputField()

// Configure the Title bar
var TitleBar *cview.TextView = cview.NewTextView()

// Box to display the operator configuration
var ActivatorBar *cview.TextView = cview.NewTextView()

// Contact input form
var ContactInputForm *cview.Form = cview.NewForm()

// Area to show logged contacts
var LogBox *cview.Box = cview.NewBox()

// Status area
var StatusBox *cview.TextView = cview.NewTextView()

// Overall Main UI
var MainUI *cview.Flex = cview.NewFlex()

func DisplayMainUI(app *cview.Application) {
	// Form element to enter hunter callsign
	WorkedCallsignInput.SetLabel("Their Call")
	WorkedCallsignInput.SetFieldWidth(10)

	// Define sent report input field
	SentReportInput.SetLabel("RST Sent")
	SentReportInput.SetFieldWidth(4)

	// Define received report input field
	RcvdReportInput.SetLabel("RST Rcvd")
	RcvdReportInput.SetFieldWidth(4)

	// Define activator frequency input field
	WorkedParkInput.SetLabel("Their Park")
	WorkedParkInput.SetFieldWidth(8)

	// Define comments input field
	CommentsInput.SetLabel("Comments")
	CommentsInput.SetFieldWidth(30)

	// Define activator frequency input field
	FreqInput.SetLabel("Freq (MHz)")
	FreqInput.SetFieldWidth(10)

	// Define activator mode in a dropdown
	ModeDropDown.SetLabel("Mode")
	ModeDropDown.SetOptionsSimple(nil,
		"CW",
		"SSB",
		"FT8",
		"FT4",
		"JS8",
	)

	// Define activator power input field
	MyPowerInput.SetLabel("Power Out")
	MyPowerInput.SetFieldWidth(4)

	// Setup the title bar
	TitleBar.SetPadding(0, 0, 0, 0)
	TitleBar.SetBorder(false)
	TitleBar.SetBackgroundColor(tcell.ColorBlue.TrueColor())
	TitleBar.SetTextColor(tcell.ColorWhite.TrueColor())
	TitleBar.SetText("mLog v0.1.0 by NQ0M")
	TitleBar.SetTextAlign(cview.AlignCenter)

	// Setup the activator bar
	ActivatorBar.SetPadding(0, 0, 0, 0)
	ActivatorBar.SetDynamicColors(true)
	ActivatorBar.SetBorder(false)
	ActivatorBar.SetText(Op.MyCallsign + "@" + Op.MyPark + " Contacts [red]" + strconv.Itoa(Op.NumContacts) + "[white]")
	ActivatorBar.SetTextAlign(cview.AlignCenter)

	// Setup the log entry form
	ContactInputForm.AddFormItem(WorkedCallsignInput)
	ContactInputForm.AddFormItem(SentReportInput)
	ContactInputForm.AddFormItem(RcvdReportInput)
	ContactInputForm.AddFormItem(WorkedParkInput)
	ContactInputForm.AddFormItem(CommentsInput)
	ContactInputForm.AddFormItem(FreqInput)
	ContactInputForm.AddFormItem(ModeDropDown)
	ContactInputForm.AddFormItem(MyPowerInput)
	ContactInputForm.AddButton("Log", func() {
		LogContact()
	})
	ContactInputForm.SetHorizontal(true)
	ContactInputForm.SetBorder(false)
	ContactInputForm.SetPadding(0, 0, 1, 1)
	ContactInputForm.SetFocus(0)

	// Create the logging box
	LogBox.SetBorder(true)
	LogBox.SetTitle("Log")
	LogBox.SetTitleAlign(cview.AlignCenter)

	// Create the status box
	StatusBox.SetBorder(true)
	StatusBox.SetTitle("Status")
	StatusBox.SetTitleAlign(cview.AlignCenter)

	MainUI.SetDirection(cview.FlexRow)
	MainUI.AddItem(TitleBar, 1, 1, false)
	MainUI.AddItem(ActivatorBar, 2, 1, false)
	MainUI.AddItem(ContactInputForm, 2, 1, true)
	MainUI.AddItem(LogBox, 0, 1, false)
	MainUI.AddItem(StatusBox, 5, 1, false)

	app.SetRoot(MainUI, true)
}
