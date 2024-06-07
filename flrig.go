package main

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"time"

	"encoding/xml"

	"code.rocketnine.space/tslocum/cview"
)

type Response struct {
	XMLName xml.Name `xml:"methodResponse"`
	Text    string   `xml:",chardata"`
	Params  struct {
		Text  string `xml:",chardata"`
		Param struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value"`
		} `xml:"param"`
	} `xml:"params"`
}

func GetFlrig(app *cview.Application) {
	for {
		rpc_req := "<?xml version='1.0'?><methodCall><methodName>rig.get_vfo</methodName></methodCall>"
		buffer := bytes.NewBufferString(rpc_req)
		rpc_resp, _ := http.Post("http://"+Op.FlrigAddress+"/", "text/xml", buffer)
		body, _ := io.ReadAll(rpc_resp.Body)
		var response Response
		err := xml.Unmarshal(body, &response)
		if err != nil {
			panic(err)
		}
		freq, _ := strconv.ParseFloat(response.Params.Param.Value, 64)
		freq = freq / 1000000
		app.QueueUpdateDraw(func() {
			FreqInput.SetText(strconv.FormatFloat(freq, 'f', 6, 64))
		})
		rpc_resp.Body.Close()
		time.Sleep(10 * time.Second)
	}
}
