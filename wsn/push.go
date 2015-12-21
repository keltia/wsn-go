// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
	"wsn-go/config"
	"wsn-ng/soap"
)

// PushClient represents an active Push mode client for WS-N.  It maintains a list of
// subscribed topics.
type PushClient struct {
	Config  *config.Config
	List    TopicList
	Timeout time.Duration
	timer   chan bool
	verbose bool
}

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
		vars := SubVars{
			TopicURL: c.createEndpoint(name),
			TopicName: name,
		}
		xmlReq, err = createRequest("subscribe", subscribePushText, vars)

		// Send SOAP request
		targetURL := c.generateURL(config.Endpoint)
		answer, err = soap.SendRequest("subscribe", targetURL, &xmlReq)

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

// Public API

// NewPushClient creates a new client using push mode with an empty list of topics.
func NewPushClient(config *config.Config) (client * PushClient) {
	return &PushClient{
		Config: config,
		List: TopicList{},
		Timeout: -1,
	}
}

// Type returns the client mode
func (c *PushClient) Type() int {
	return MODE_PUSH
}

// SetVerbose is obvious
func (c *PushClient) SetVerbose() {
	c.verbose = true
}

// Subscribe registers the given topic
func (c *PushClient) Subscribe(topic string) (err error) {
	// Add the topic
	log.Printf("subscribe push/%s", topic)
	if _, ok := c.List[topic]; ok {
		err = ErrTopicAlreadySubscribed
	}
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
		err = c.realUnsubscribe(name)
	} else {
		err = ErrTopicNotFound
	}
	return
}

// SetTimeout records that we want to stop after some time
func (c *PushClient) SetTimeout(timeout int64) {
	c.Timeout = timeout * time.Second
}

// Start does the real subscribe because it actually start the data flow
func (c *PushClient) Start() (err error) {
	// Setup callback server

	// Setup the subscriptions, data will flow now
	for name, _ := range c.List {
		if err = c.realSubscribe(name); err != nil {
			return
		}
	}

	// Set timer
	if c.Timeout != -1 {
		var expired bool

		c.timer = make(chan bool)
		go func() {
			time.Sleep(time.Duration(c.Timeout) * time.Second)
			log.Println("Timer expired!")
			c.timer <- true
		}()
		// Wait for timeout
		expired <- c.timer
		if expired {
			c.Stop()
		} else {
			log.Fatalf("Fatal: timer expired for unknown reason")
		}
	}
	return
}

// Stop does unsubscribe all the topics at once
func (c *PushClient) Stop() (err error) {
	log.Printf("stopping everything")

	for name, _ := range c.List {
		err = c.Unsubscribe(name)
	}
	// Stop callback server
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
