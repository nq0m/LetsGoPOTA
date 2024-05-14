package main

import (
	"code.rocketnine.space/tslocum/cview"
)

// Define all UI elements as global variables
var Op Operator

func main() {

	app := cview.NewApplication()
	app.EnableMouse(true)

	//Run the config UI first
	DisplayConfigUI(app)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
