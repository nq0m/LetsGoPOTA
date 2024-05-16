package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "modernc.org/sqlite"
)

func CreateLogFile() {
	//Get the current user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dotconigdir := home + "/.config"
	// Check if our ~/.config directory exists
	if _, err := os.Stat(dotconigdir); os.IsNotExist(err) {
		// If it doesn't exist, create it
		err := os.Mkdir(dotconigdir, 0755)
		if err != nil {
			panic(err)
		}
	}
	// Check if our ~/.config/mlog directory exists
	if _, err := os.Stat(dotconigdir + "/mlog"); os.IsNotExist(err) {
		// If it doesn't exist, create it
		err := os.Mkdir(dotconigdir+"/mlog", 0755)
		if err != nil {
			panic(err)
		}
	}
	// Define our database file path
	dbFile := dotconigdir + "/mlog/" + Op.MyCallsign + "@" + Op.MyPark + "-" + time.Now().UTC().Format("20060102") + ".db"
	// Open our new database file
	Op.Database, err = sql.Open("sqlite", dbFile)
	if err != nil {
		panic(err)
	}
	if _, err = Op.Database.Exec(`
CREATE TABLE IF NOT EXISTS log (
call TEXT NOT NULL,
band TEXT NOT NULL,
mode TEXT NOT NULL,
qso_date TEXT NOT NULL,
time_on TEXT NOT NULL,
freq TEXT NOT NULL,
rst_sent TEXT,
rst_rcvd TEXT,
sig_info TEXT,
comment TEXT,
tx_pwr TEXT
);
	`); err != nil {
		panic(err)
	}
	// Prepared statement to insert a contact into the database
	Op.LogStatement, err = Op.Database.Prepare("INSERT INTO log (call, band, mode, qso_date, time_on, freq, rst_sent, rst_rcvd, sig_info, comment, tx_pwr) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
}

func LogContact() {
	// First, we need to get our mode from the dropdown options
	_, selectedOption := ModeDropDown.GetCurrentOption()
	mode := selectedOption.GetText()
	// Also we need to get our band based on our frequency
	band := get_band_from_freq(FreqInput.GetText())
	// Compose our query
	dbQuery := fmt.Sprintf(
		`INSERT INTO log (
call,
band,
mode,
qso_date,
time_on,
freq,
rst_sent,
rst_rcvd,
sig_info,
comment,
tx_pwr
) VALUES (
'%s',
'%s',
'%s',
'%s',
'%s',
'%s',
'%s',
'%s',
'%s',
'%s',
'%s'
)
`,
		WorkedCallsignInput.GetText(),
		band,
		mode,
		time.Now().UTC().Format("20060102"),
		time.Now().UTC().Format("15:04:05"),
		FreqInput.GetText(),
		SentReportInput.GetText(),
		RcvdReportInput.GetText(),
		WorkedParkInput.GetText(),
		CommentsInput.GetText(),
		MyPowerInput.GetText(),
	)
	// TEST: dump our query to file
	err := os.WriteFile("/tmp/query.txt", []byte(dbQuery), 0644)
	if err != nil {
		panic(err)
	}
	// Execute the prepared query with our values
	_, err = Op.LogStatement.Exec(
		WorkedCallsignInput.GetText(),
		band,
		mode,
		time.Now().UTC().Format("20060102"),
		time.Now().UTC().Format("15:04:05"),
		FreqInput.GetText(),
		SentReportInput.GetText(),
		RcvdReportInput.GetText(),
		WorkedParkInput.GetText(),
		CommentsInput.GetText(),
		MyPowerInput.GetText(),
	)
	if err != nil {
		panic(err)
	}
	// Now we need to blank some of the fields
	WorkedCallsignInput.SetText("")
	SentReportInput.SetText("")
	RcvdReportInput.SetText("")
	WorkedParkInput.SetText("")
	CommentsInput.SetText("")
	// Increment the number of contacts
	Op.NumContacts = Op.NumContacts + 1
	// and update the title bar with the number of contacts
	if Op.NumContacts >= 10 {
		// 10 or more contacts, successful activation, make the number green
		ActivatorBar.SetText(Op.MyCallsign + "@" + Op.MyPark + " Contacts [green]" + strconv.Itoa(Op.NumContacts) + "[white]")
	} else {
		// Less than 10 contacts, keep the number in red text
		ActivatorBar.SetText(Op.MyCallsign + "@" + Op.MyPark + " Contacts [red]" + strconv.Itoa(Op.NumContacts) + "[white]")
	}
	// and set focus back to the callsign field
	ContactInputForm.SetFocus(0)
}

func get_band_from_freq(freq string) string {
	// Convert the frequency to a float
	f, err := strconv.ParseFloat(freq, 32)
	if err != nil {
		panic(err)
	}
	if f >= 1.800 && f <= 2.000 {
		return "160m"
	} else if f >= 3.500 && f <= 4.000 {
		return "80m"
	} else if f >= 5.3305 && f <= 5.4065 {
		return "60m"
	} else if f >= 7.000 && f <= 7.300 {
		return "40m"
	} else if f >= 10.100 && f <= 10.150 {
		return "30m"
	} else if f >= 14.000 && f <= 14.350 {
		return "20m"
	} else if f >= 18.068 && f <= 18.168 {
		return "17m"
	} else if f >= 21.000 && f <= 21.450 {
		return "15m"
	} else if f >= 24.890 && f <= 24.990 {
		return "12m"
	} else if f >= 28.000 && f <= 29.700 {
		return "10m"
	} else if f >= 50.000 && f <= 54.000 {
		return "6m"
	} else if f >= 144.000 && f <= 148.000 {
		return "2m"
	} else if f >= 222.000 && f <= 225.000 {
		return "1.25m"
	} else if f >= 420.000 && f <= 450.000 {
		return "70cm"
	} else {
		return "UNKNOWN"
	}
}