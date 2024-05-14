package main

import (
	"strings"

	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var MyCallsignInput *cview.InputField = cview.NewInputField()
var MyGridInput *cview.InputField = cview.NewInputField()
var MyStateInput *cview.DropDown = cview.NewDropDown()
var MyParkInput *cview.InputField = cview.NewInputField()

func convert_input_to_upper(input *cview.InputField) {
	input_text := strings.ToUpper(input.GetText())
	input.SetText(input_text)
}

func DisplayConfigUI(app *cview.Application) {
	// Form element to enter activator callsign
	MyCallsignInput.SetLabel("My Call")
	MyCallsignInput.SetFieldWidth(10)
	MyCallsignInput.SetDoneFunc(func(key tcell.Key) {
		convert_input_to_upper(MyCallsignInput)
	})

	// Form Element to enter activator grid square
	MyGridInput.SetLabel("My Grid Square")
	MyGridInput.SetFieldWidth(5)
	MyGridInput.SetDoneFunc(func(key tcell.Key) {
		convert_input_to_upper(MyGridInput)
	})

	// Form element to enter activator state
	MyStateInput.SetLabel("My State")
	MyStateInput.SetOptionsSimple(nil,
		"AL",
		"AK",
		"AZ",
		"AR",
		"CA",
		"CO",
		"CT",
		"DC",
		"DE",
		"FL",
		"GA",
		"HI",
		"ID",
		"IL",
		"IN",
		"IA",
		"KS",
		"KY",
		"LA",
		"ME",
		"MD",
		"MA",
		"MI",
		"MN",
		"MS",
		"MO",
		"MT",
		"NE",
		"NV",
		"NH",
		"NJ",
		"NM",
		"NY",
		"NC",
		"ND",
		"OH",
		"OK",
		"OR",
		"PA",
		"RI",
		"SC",
		"SD",
		"TN",
		"TX",
		"UT",
		"VT",
		"VA",
		"WA",
		"WV",
		"WI",
		"WY",
	)

	// Form Element to enter activator grid square
	MyParkInput.SetLabel("My Park")
	MyParkInput.SetFieldWidth(9)
	MyParkInput.SetDoneFunc(func(key tcell.Key) {
		convert_input_to_upper(MyParkInput)
	})

	// Create the config form
	ConfigForm := cview.NewForm()
	ConfigForm.AddFormItem(MyCallsignInput)
	ConfigForm.AddFormItem(MyGridInput)
	ConfigForm.AddFormItem(MyStateInput)
	ConfigForm.AddFormItem(MyParkInput)
	ConfigForm.AddButton("Finished", func() {
		StoreConfiguration()
		CreateLogFile()
		DisplayMainUI(app)
	})
	ConfigForm.SetFocus(0)

	app.SetRoot(ConfigForm, true)
}

func StoreConfiguration() {
	// Store operator config
	Op.MyCallsign = MyCallsignInput.GetText()
	Op.MyGridSquare = MyGridInput.GetText()
	_, selectedOption := MyStateInput.GetCurrentOption()
	Op.MyState = selectedOption.GetText()
	Op.MyPark = MyParkInput.GetText()
	Op.NumContacts = 0
}
