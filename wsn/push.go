// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"strings"
	"wsn-go/config"
	"wsn-ng/soap"
)

// A PullClient represents an active Push mode client for WS-N.  It maintains a list of
// subscribed topics.
type PushClient struct {
	Config *config.Config
	List TopicList
}

// Private func

// createEndpoint generates our local endpoint URL
func (c *PushClient) createEndpoint(name string) (endpoint string) {
	config := c.Config
	realEP := strings.ToLower(name)
	endpoint = fmt.Sprintf("%s:%d%s", config.Base, config.Port, realEP)
	endpoint = fmt.Sprintf("%s://%s:%d/%s", config.Proto, config.Site, config.Port, realEP)
	return
}

// Generate an URL on the target site
func (c *PushClient) generateURL(endPoint string) string {
	config := c.Config
	return fmt.Sprintf("%s://%s:%d/%s", config.Proto, config.Site, config.Port, endPoint)
}

// Does the actual WS-N subscription
func (c *PushClient) realSubscribe(name string) (err error) {
	var config = c.Config

	// Handle only already registered topics
	if _, present := c.List[name]; present {
		var xmlReq bytes.Buffer
		var body []byte

		// Prepare the request
		vars := SubVars{
			TopicURL: c.createEndpoint(name),
			TopicName: name,
		}
		xmlReq, err = createRequest("subscribe", subscribePushText, vars)

		// Send SOAP request
		targetURL := c.generateURL(config.Endpoint)
		body, err = soap.SendRequest(targetURL, xmlReq)

		// Parse XML
		res := &SubscribeAnswer{}
		if err = xml.Unmarshal(body, res); err != nil {
			return
		}

		// Special case
		address := res.Body.Resp.Reference.Address
		address = strings.Replace(address, "0.0.0.0", config.Site, -1)

		c.List[name].UnsubAddr = address
		c.List[name].Started = true
	} else {
		err = ErrTopicNotFound
	}
	return
}

// Does the actual WS-N un-subscription
func (c *PushClient) realUnsubscribe(name string) (err error) {
	if _, present := c.List[name]; present {
		c.List[name].UnsubAddr = ""
		c.List[name].Started = false
	} else {
		err = ErrTopicNotFound
	}
	return
}

// Public API

// NewPushClient creates a new client using push mode with an empty list of topics.
func NewPushClient(config *config.Config) (client * PushClient) {
	return &PushClient{
		Config: config,
		List: TopicList{},
	}
}

// Type returns the client mode
func (c *PushClient) Type() int {
	return MODE_PUSH
}

// Subscribe registers the given topic
func (c *PushClient) Subscribe(topic string) (err error) {
	// Add the topic
	log.Printf("subscribe push/%s", topic)
	c.List[topic] = &Topic{
		Started: false,
		UnsubAddr: "",
		Bytes: 0,
		Pkts: 0,
	}

	return
}

// Unsubscribe stops and closes the given topic
func (c *PushClient) Unsubscribe(name string) (err error) {
	log.Printf("unsubscribed push/%s", name)

	if _, present := c.List[name]; present {
		// xml.doUnsubscribe()
		c.List[name].Started = false
	} else {
		err = ErrTopicNotFound
	}
	return
}

// Start does the real subscribe because it actually start the data flow
func (c *PushClient) Start() (err error) {
	for name, _ := range c.List {
		err = c.realSubscribe(name)
	}
	return
}

// Stop does unsubscribe all the topics at once
func (c *PushClient) Stop() (err error) {
	log.Printf("stopping everything")

	for name, _ := range c.List {
		err = c.Unsubscribe(name)
	}
	return
}

// Read is needed for the io.Reader interface
func (c *PushClient) Read(p []byte) (n int, err error) {
	data := "push/foobar"
	n = len(data)
	copy(p, []byte(data))
	err = io.EOF
	return
}
