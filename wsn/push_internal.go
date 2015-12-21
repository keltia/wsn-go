// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
    "encoding/xml"
    "fmt"
    "strings"

    "wsn-go/soap"
)

// Private func

// createEndpoint generates our local endpoint URL
func (c *PushClient) createEndpoint(name string) (endpoint string) {
	realEP := "/" + strings.ToLower(name)
	endpoint = fmt.Sprintf("%s:%d%s", c.base, c.port, realEP)
	return
}

// generateURL generates an URL on the target site
func (c *PushClient) generateURL(endPoint string) string {
    return fmt.Sprintf("%s:%d/%s", c.target, c.port, endPoint)
}

// realSubscribe does the actual WS-N subscription
func (c *PushClient) realSubscribe(name string) (err error) {

	// Handle only already registered topics
	if topic, present := c.List[name]; present {
		var answer []byte

		// Prepare the request
		vars := soap.SubVars{
			TopicURL: c.createEndpoint(name),
			TopicName: name,
		}
		soapReq, err := soap.NewRequest(soap.SUBSCRIBEPUSH, vars)
	        if err != nil {
		    return
		}

		// Send SOAP request
		answer, err = soapReq.Send(c.target)
	    	if err != nil {
		    return
		}

	    	// We will fix the broken return address (might be 0.0.0.0)
	    	baseIp := strings.Split(c.target, ":")
		address := strings.Replace(answer, "0.0.0.0", baseIp, -1)
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
            // Prepare the request
	    soapReq, err := soap.NewRequest(soap.UNSUBSCRIBEPUSH, soap.SubVars{})
	    if err != nil {
		return
	    }
	    // Send SOAP request
	    _, err = soapReq.Send(topic.UnsubAddr)

            topic.UnsubAddr = ""
            topic.Started = false
	} else {
		err = ErrTopicNotFound
	}
	return
}

