package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"unicode/utf8"

	_ "modernc.org/sqlite"
)

func ExportADIF() {
	// Setup our file path & name
	adifFilename := GetDocumentsDirectory() + Op.MyCallsign + "@" + Op.MyPark + "-" + time.Now().UTC().Format("20060102") + ".adif"
	// Output info to status box
	StatusBox.SetText("Writing ADIF file to: " + adifFilename)
	// Create our file
	adifFile, err := os.Create(adifFilename)
	if err != nil {
		panic(err)
	}
	defer adifFile.Close()
	// Write our header
	fmt.Fprintf(adifFile,
		`Generated on %s
<adif_ver:5>3.1.4
<programid:10>LetsGoPOTA
<programversion:5>%s
<EOH>

`,
		time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		Version,
	)
	// Get data from the sqlite database
	rows, err := Op.Database.Query("SELECT * FROM log")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		qsoRecord := ""
		var call, band, mode, qso_date, time_on, freq, rst_sent, rst_rcvd, sig_info, comment, tx_pwr string
		if err := rows.Scan(&call, &band, &mode, &qso_date, &time_on, &freq, &rst_sent, &rst_rcvd, &sig_info, &comment, &tx_pwr); err != nil {
			panic(err)
		}
		qsoRecord += "<band:" + strconv.Itoa(utf8.RuneCountInString(band)) + ">" + band + "\n"
		qsoRecord += "<call:" + strconv.Itoa(utf8.RuneCountInString(call)) + ">" + call + "\n"
		qsoRecord += "<freq:" + strconv.Itoa(utf8.RuneCountInString(freq)) + ">" + freq + "\n"
		qsoRecord += "<mode:" + strconv.Itoa(utf8.RuneCountInString(mode)) + ">" + mode + "\n"
		qsoRecord += "<my_gridsquare:" + strconv.Itoa(utf8.RuneCountInString(Op.MyGridSquare)) + ">" + Op.MyGridSquare + "\n"
		qsoRecord += "<my_sig:4>POTA\n"
		qsoRecord += "<my_sig_info:" + strconv.Itoa(utf8.RuneCountInString(Op.MyPark)) + ">" + Op.MyPark + "\n"
		qsoRecord += "<my_state:" + strconv.Itoa(utf8.RuneCountInString(Op.MyState)) + ">" + Op.MyState + "\n"
		qsoRecord += "<operator:" + strconv.Itoa(utf8.RuneCountInString(Op.MyCallsign)) + ">" + Op.MyCallsign + "\n"
		qsoRecord += "<qso_date:" + strconv.Itoa(utf8.RuneCountInString(qso_date)) + ">" + qso_date + "\n"
		qsoRecord += "<time_on:" + strconv.Itoa(utf8.RuneCountInString(time_on)) + ">" + time_on + "\n"
		if rst_rcvd != "" {
			qsoRecord += "<rst_rcvd:" + strconv.Itoa(utf8.RuneCountInString(rst_rcvd)) + ">" + rst_rcvd + "\n"
		}
		if rst_sent != "" {
			qsoRecord += "<rst_sent:" + strconv.Itoa(utf8.RuneCountInString(rst_sent)) + ">" + rst_sent + "\n"
		}
		if sig_info != "" {
			qsoRecord += "<sig_info:" + strconv.Itoa(utf8.RuneCountInString(sig_info)) + ">" + sig_info + "\n"
		}
		if comment != "" {
			qsoRecord += "<comment:" + strconv.Itoa(utf8.RuneCountInString(comment)) + ">" + comment + "\n"
		}
		if tx_pwr != "" {
			qsoRecord += "<tx_pwr:" + strconv.Itoa(utf8.RuneCountInString(tx_pwr)) + ">" + tx_pwr + "\n"
		}
		qsoRecord += "<eor>\n\n"
		fmt.Fprint(adifFile, qsoRecord)
	}
}
