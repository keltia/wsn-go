// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
    "bytes"
    "encoding/xml"
    "fmt"
    "strings"

    "wsn-go/soap"
)

// Private func

// createEndpoint generates our local endpoint URL
func (c *PushClient) createEndpoint(name string) (endpoint string) {
	config := c.Config
	realEP := "/" + strings.ToLower(name)
	endpoint = fmt.Sprintf("%s:%d%s", config.Base, config.Port, realEP)
	return
}

// generateURL generates an URL on the target site
func (c *PushClient) generateURL(endPoint string) string {
	config := c.Config
	return fmt.Sprintf("%s://%s:%d/%s", config.Proto, config.Site, config.Port, endPoint)
}

// realSubscribe does the actual WS-N subscription
func (c *PushClient) realSubscribe(name string) (err error) {
	var config = c.Config

	// Handle only already registered topics
	if topic, present := c.List[name]; present {
		var xmlReq bytes.Buffer
		var answer []byte

		// Prepare the request
		vars := soap.SubVars{
			TopicURL: c.createEndpoint(name),
			TopicName: name,
		}
		soapReq, err := soap.NewRequest(soap.SUBSCRIBEPUSH, vars)

		// Send SOAP request
		targetURL := c.generateURL(config.Endpoint)
		answer, err = soapReq.Send(targetURL)

		// Parse XML
		res := &SubscribeAnswer{}
		if err = xml.Unmarshal(answer, res); err != nil {
			return
		}

		// Special case
		address := res.Body.Resp.Reference.Address
		address = strings.Replace(address, "0.0.0.0", config.Site, -1)

		topic.UnsubAddr = address
		topic.Started = true
	} else {
		err = ErrTopicNotFound
	}
	return
}

// realUnsubscribe does the actual WS-N un-subscription
func (c *PushClient) realUnsubscribe(name string) (err error) {
	if topic, present := c.List[name]; present {
		var xmlReq *bytes.Buffer

		// Prepare the request
		xmlReq = bytes.NewBufferString(unsubscribePushText)

		// Send SOAP request
		targetURL := topic.UnsubAddr
		_, err = soap.SendRequest("Unsubscribe", targetURL, xmlReq)

		topic.UnsubAddr = ""
		topic.Started = false
	} else {
		err = ErrTopicNotFound
	}
	return
}

