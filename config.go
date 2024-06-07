package main

import "database/sql"

type Operator struct {
	MyCallsign   string
	MyGridSquare string
	MyState      string
	MyPark       string
	Database     *sql.DB
	DatabaseFile string
	LogStatement *sql.Stmt
	NumContacts  int
	FlrigAddress string
	QrzEnabled   bool
	QrzUsername  string
	QrzPassword  string
}

const Version string = "0.1.0"
const Title string = "LetsGoPota v" + Version + "by NQ0M"
