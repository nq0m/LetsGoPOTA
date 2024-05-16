package main

import "database/sql"

type Operator struct {
	MyCallsign   string
	MyGridSquare string
	MyState      string
	MyPark       string
	Database     *sql.DB
	LogStatement *sql.Stmt
	NumContacts  int
}

const Version string = "0.1.0"
const TitleText string = "LetsGoPota v" + Version + "by NQ0M"