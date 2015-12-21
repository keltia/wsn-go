// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package soap

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	httpClient http.Client = http.Client{}
)

func SendRequest(action int, targetURL string, request *bytes.Buffer) (body []byte, err error) {

	// Prepare the request
	buf := bytes.NewBufferString(request.String())
	req, err := http.NewRequest("POST", targetURL, buf)
	if err != nil {
		log.Fatal("Error creating request for ", buf, ": ", err)
	}
	req.Header.Set("SOAPAction", actionToHeader[action])
	req.Header.Set("Content-Type", "text/xml; charset=UTF-8")

	resp, err := httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		body = nil
	} else {
		// body is the XML encoded answer, to be decoded further up
		body, err = ioutil.ReadAll(resp.Body)
	}
	return
}