// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package soap

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	httpClient http.Client = http.Client{}
)

type SoapAnswer struct {
	XMLName xml.Name
	Body    []byte `xml:",innerxml"`
}

// Send the prepared request to the target SOAP endpoint
func (request *Request) Send(targetURL string) (res []byte, err error) {

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
		res = []byte{}
	} else {
		// Everything is generic at this point
		var body []byte

		// body is the XML encoded answer, to be decoded further up
		body, err = ioutil.ReadAll(resp.Body)

		// Parse XML
		answer := &SoapAnswer{}
		if err = xml.Unmarshal(body, answer); err != nil {
			return
		}
		// Now in res we have to decode the Body itself

	}
	return
}
