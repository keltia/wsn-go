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

func SendRequest(targetURL string, buf bytes.Buffer) (err error) {

	// Prepare the request
	req, err := http.NewRequest("POST", targetURL, buf)
	if err != nil {
		log.Fatal("Error creating request for ", buf, ": ", err)
	}
	req.Header.Set("SOAPAction", "Subscribe")
	req.Header.Set("Content-Type", "text/xml; charset=UTF-8")

	resp, err := httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.Printf("Error during POST: %v", err)
		return "", nil
	}

	body, err := ioutil.ReadAll(resp.Body)

}