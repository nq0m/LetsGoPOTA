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
