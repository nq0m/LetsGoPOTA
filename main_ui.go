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
var LogTable *cview.Table = cview.NewTable()

// Status area
var StatusBox *cview.TextView = cview.NewTextView()

// Overall Main UI
var MainUI *cview.Flex = cview.NewFlex()

func DisplayMainUI(app *cview.Application) {
	// Form element to enter hunter callsign
	WorkedCallsignInput.SetLabel("Their Call")
	WorkedCallsignInput.SetFieldWidth(10)
	WorkedCallsignInput.SetDoneFunc(func(key tcell.Key) {
		convert_input_to_upper(WorkedCallsignInput)
		go QrzLookup(app)
	})

	// Define sent report input field
	SentReportInput.SetLabel("RST Sent")
	SentReportInput.SetFieldWidth(4)

	// Define received report input field
	RcvdReportInput.SetLabel("RST Rcvd")
	RcvdReportInput.SetFieldWidth(4)

	// Define activator frequency input field
	WorkedParkInput.SetLabel("Their Park")
	WorkedParkInput.SetFieldWidth(9)
	WorkedParkInput.SetDoneFunc(func(key tcell.Key) {
		convert_input_to_upper(WorkedParkInput)
	})

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

	// Create our full TitleText, including status indicators
	TitleText := Title
	if Op.FlrigAddress != "" {
		TitleText = TitleText + " [green]FLRig[white] "
	}
	if Op.QrzEnabled {
		TitleText = TitleText + " [green]QRZ[white] "
	}

	// Setup the title bar
	TitleBar.SetPadding(0, 0, 0, 0)
	TitleBar.SetBorder(false)
	TitleBar.SetDynamicColors(true)
	TitleBar.SetBackgroundColor(tcell.ColorBlue.TrueColor())
	TitleBar.SetTextColor(tcell.ColorWhite.TrueColor())
	TitleBar.SetText(TitleText)
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
		LogContact(app)
		DisplayLogUIContacts(app)
	})
	ContactInputForm.AddButton("Export ADIF", func() {
		ExportADIF()
	})
	ContactInputForm.SetHorizontal(true)
	ContactInputForm.SetBorder(false)
	ContactInputForm.SetPadding(0, 0, 1, 1)
	ContactInputForm.SetFocus(0)

	DisplayLogUIHeadings(app)
	// If we have existing contacts, we need to display them
	if Op.NumContacts > 0 {
		DisplayLogUIContacts(app)
	}

	// Create the status box
	StatusBox.SetBorder(true)
	StatusBox.SetTitle("Status")
	StatusBox.SetTitleAlign(cview.AlignCenter)

	MainUI.SetDirection(cview.FlexRow)
	MainUI.AddItem(TitleBar, 1, 1, false)
	MainUI.AddItem(ActivatorBar, 2, 1, false)
	MainUI.AddItem(ContactInputForm, 2, 1, true)
	MainUI.AddItem(LogTable, 0, 1, false)
	MainUI.AddItem(StatusBox, 5, 1, false)

	app.SetRoot(MainUI, true)
	if Op.FlrigAddress != "" {
		go GetFlrig(app)
	}
}
