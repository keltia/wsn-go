// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package soap

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
"encoding/xml"
"strings"
"wsn-go/config"
)

var (
	httpClient http.Client = http.Client{}
)

type SubscribeAnswer struct {
	XMLName xml.Name
	Body    SABody
}

type SABody struct {
	XMLName xml.Name
	Resp    SAResp `xml:"SubscribeResponse"`
}

type SAResp struct {
	XMLName   xml.Name    `xml:"SubscribeResponse"`
	Reference SAReference `xml:"SubscriptionReference"`
}

type SAReference struct {
	XMLName             xml.Name `xml:"SubscriptionReference"`
	Address             string
	ReferenceParameters string
	Metadata            string
}

// Send the prepared request to the target SOAP endpoint
func (request *Request) Send(targetURL string) (address string, err error) {

	// Prepare the request
	buf := bytes.NewBufferString(request.Text.String())
	req, err := http.NewRequest("POST", targetURL, buf)
	if err != nil {
		log.Fatal("Error creating request for ", buf, ": ", err)
	}
	req.Header.Set("SOAPAction", actionToHeader[request.Action])
	req.Header.Set("Content-Type", "text/xml; charset=UTF-8")

	resp, err := httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		address = nil
	} else {
		// body is the XML encoded answer, to be decoded further up
		body, err := ioutil.ReadAll(resp.Body)

		// Parse XML
		res := &SubscribeAnswer{}
		if err = xml.Unmarshal(body, res); err != nil {
			return
		}
		address = res.Body.Resp.Reference.Address
	}
	return
}