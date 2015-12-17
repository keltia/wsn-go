// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"io"
	"log"
)

// A PullClient represents an active Push mode client for WS-N.  It maintains a list of
// subscribed topics.
type PushClient struct {
	List TopicList
}

// Private func

// Does the actual WS-N subscription
func (c *PushClient) realSubscribe(name string) (err error) {
	if _, present := c.List[name]; present {
		c.List[name].UnsubAddr = "my_unsub_addr"
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
func NewPushClient() (client * PushClient) {
	return &PushClient{
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
