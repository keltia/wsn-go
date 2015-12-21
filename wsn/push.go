// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"io"
	"log"
	"net/http"
	"time"

	"wsn-go/config"
)

// PushClient represents an active Push mode client for WS-N.  It maintains a list of
// subscribed topics.
type PushClient struct {
	Config  *config.Config
	List    TopicList
	Timeout time.Duration

	// Private fields
	base    string
	target  string
	port    int
	server  http.Server
	verbose bool
}

// Public API

// NewPushClient creates a new client using push mode with an empty list of topics.
func NewPushClient(config *config.Config) (client * PushClient) {
	client = &PushClient{
		Config: config,
		List: TopicList{},
		Timeout: -1,
		base: config.Base,
		target: config.Target,
		port: config.Port,
		verbose: false,
	}
	client.target = client.generateURL(config.Broker)
	return
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
		// Fire up goroutine
		go func() {
			time.Sleep(time.Duration(c.Timeout) * time.Second)
			log.Println("Timer expired!")
			c.Stop()
		}()
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
