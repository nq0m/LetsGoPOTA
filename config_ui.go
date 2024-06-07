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
var FlrigInput *cview.InputField = cview.NewInputField()
var QrzCheckBox *cview.CheckBox = cview.NewCheckBox()
var QrzUsername *cview.InputField = cview.NewInputField()
var QrzPassword *cview.InputField = cview.NewInputField()
var FinishButton *cview.Button = cview.NewButton("Finished")
var emptyTextView *cview.TextView = cview.NewTextView()

var ConfigUI *cview.Grid = cview.NewGrid()

func convert_input_to_upper(input *cview.InputField) {
	input_text := strings.ToUpper(input.GetText())
	input.SetText(input_text)
}

func DisplayConfigUI(app *cview.Application) {
	// Form element to enter activator callsign
	MyCallsignInput.SetLabel("My Call: ")
	MyCallsignInput.SetFieldWidth(10)
	MyCallsignInput.SetDoneFunc(func(key tcell.Key) {
		convert_input_to_upper(MyCallsignInput)
		if key == tcell.KeyTab {
			app.SetFocus(MyGridInput)
		}
	})

	// Form Element to enter activator grid square
	MyGridInput.SetLabel("My Grid Square: ")
	MyGridInput.SetFieldWidth(5)
	MyGridInput.SetDoneFunc(func(key tcell.Key) {
		convert_input_to_upper(MyGridInput)
		if key == tcell.KeyTab {
			app.SetFocus(MyStateInput)
		}
		if key == tcell.KeyBacktab {
			app.SetFocus(MyCallsignInput)
		}
	})

	// Form element to enter activator state
	MyStateInput.SetLabel("My State: ")
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
	MyStateInput.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyTab {
			app.SetFocus(MyParkInput)
		}
		if key == tcell.KeyBacktab {
			app.SetFocus(MyGridInput)
		}
	})

	// Form Element to enter activator grid square
	MyParkInput.SetLabel("My Park: ")
	MyParkInput.SetFieldWidth(9)
	MyParkInput.SetDoneFunc(func(key tcell.Key) {
		convert_input_to_upper(MyParkInput)
		if key == tcell.KeyTab {
			app.SetFocus(FlrigInput)
		}
		if key == tcell.KeyBacktab {
			app.SetFocus(MyStateInput)
		}
	})

	// Form element to enter FLRig address
	FlrigInput.SetLabel("Optional FLRig Address (ip:port): ")
	FlrigInput.SetFieldWidth(21)
	FlrigInput.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyTab {
			app.SetFocus(QrzCheckBox)
		}
		if key == tcell.KeyBacktab {
			app.SetFocus(MyParkInput)
		}
	})

	QrzCheckBox.SetLabel("Use QRZ API: ")
	QrzCheckBox.SetChecked(false)
	QrzCheckBox.SetChangedFunc(func(checked bool) {
		if checked {
			ConfigUI.RemoveItem(QrzCheckBox)
			ConfigUI.AddItem(QrzCheckBox, 2, 0, 1, 1, 0, 0, true)
			ConfigUI.AddItem(QrzUsername, 2, 1, 1, 1, 0, 0, false)
			ConfigUI.AddItem(QrzPassword, 2, 2, 1, 2, 0, 0, false)
		} else {
			ConfigUI.RemoveItem(QrzCheckBox)
			ConfigUI.RemoveItem(QrzUsername)
			ConfigUI.RemoveItem(QrzPassword)
			ConfigUI.AddItem(QrzCheckBox, 2, 0, 1, 4, 0, 0, true)
		}
	})
	QrzCheckBox.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyTab {
			if QrzCheckBox.IsChecked() {
				app.SetFocus(QrzUsername)
			}
		} else {
			app.SetFocus(FinishButton)
		}
		if key == tcell.KeyBacktab {
			app.SetFocus(FlrigInput)
		}
	})

	QrzUsername.SetLabel("QRZ Username: ")
	QrzUsername.SetFieldWidth(10)
	QrzUsername.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyTab {
			app.SetFocus(QrzPassword)
		}
		if key == tcell.KeyBacktab {
			app.SetFocus(QrzCheckBox)
		}
	})

	QrzPassword.SetLabel("QRZ Password: ")
	QrzPassword.SetFieldWidth(10)
	QrzPassword.SetMaskCharacter('*')
	QrzPassword.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyTab {
			app.SetFocus(FinishButton)
		}
		if key == tcell.KeyBacktab {
			app.SetFocus(QrzUsername)
		}
	})

	emptyTextView.SetText("")
	emptyTextView.SetScrollBarVisibility(cview.ScrollBarNever)

	FinishButton.SetLabel("Finished")
	FinishButton.SetSelectedFunc(func() {
		StoreConfiguration()
		CreateLogFile()
		DisplayMainUI(app)
	})

	ConfigUI.SetColumns(0, 0, 0, 0)
	ConfigUI.SetRows(3, 3, 3, 3, 0)
	ConfigUI.AddItem(MyCallsignInput, 0, 0, 1, 1, 0, 0, true)
	ConfigUI.AddItem(MyGridInput, 0, 1, 1, 1, 0, 0, false)
	ConfigUI.AddItem(MyStateInput, 0, 2, 1, 1, 0, 0, false)
	ConfigUI.AddItem(MyParkInput, 0, 3, 1, 1, 0, 0, false)
	ConfigUI.AddItem(FlrigInput, 1, 0, 1, 4, 0, 0, false)
	ConfigUI.AddItem(QrzCheckBox, 2, 0, 1, 4, 0, 0, false)
	ConfigUI.AddItem(FinishButton, 3, 0, 1, 4, 0, 0, false)
	ConfigUI.AddItem(emptyTextView, 4, 0, 1, 4, 0, 0, false)

	app.SetRoot(ConfigUI, true)
}

func StoreConfiguration() {
	// Store operator config
	Op.MyCallsign = MyCallsignInput.GetText()
	Op.MyGridSquare = MyGridInput.GetText()
	_, selectedOption := MyStateInput.GetCurrentOption()
	Op.MyState = selectedOption.GetText()
	Op.MyPark = MyParkInput.GetText()
	Op.FlrigAddress = FlrigInput.GetText()
	if QrzCheckBox.IsChecked() {
		Op.QrzEnabled = true
		Op.QrzUsername = QrzUsername.GetText()
		Op.QrzPassword = QrzPassword.GetText()
	} else {
		Op.QrzEnabled = false
	}
	Op.NumContacts = 0
}
