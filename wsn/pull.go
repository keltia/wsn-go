// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"io"
	"log"
	"wsn-go/config"
)

// A PullClient represents an active Pull mode client for WS-N.  It maintains a list of
// subscribed topics and the Pull point that will be used to subscribe.
type PullClient struct {
	PullPt string
	List   TopicList
	Timeout int64

	// Private fields
	base    string
	target  string
	port    int
	verbose bool
	output  chan []byte
}

// NewPullClient creates a new instance of a Pull client.
func NewPullClient(config *config.Config) (client *PullClient) {
	client = &PullClient{
		PullPt: "",
		List:   TopicList{},
		Timeout: -1,
		base:    config.Base,
		target:  config.Target,
		port:    config.Port,
		verbose: false,
	}
	client.target = client.generateURL(config.Broker)
	return
}

// Type returns the operating mode of the client
func (c *PullClient) Type() int {
	return MODE_PULL
}

// Subscribe register a topic for future consumption.  It also create the Pull point on first
// use.
func (c *PullClient) Subscribe(topic string) (err error) {

	// Create Pull Point if needed
	if c.PullPt == "" {
		log.Printf("creating pull point for %s", topic)
		if c.PullPt, err = createPullPoint(); err != nil {
			return
		}
	}

	// Add the topic
	log.Printf("subscribe pull/%s", topic)
	c.List[topic] = &Topic{
		Started:   false,
		UnsubAddr: "",
		Bytes:     0,
		Pkts:      0,
	}

	return
}

// Unsubscribe unregister the topic from the list of active ones.  It also destroys the Pull
// point on last Unsubscribe.
func (c *PullClient) Unsubscribe(topic string) (err error) {
	log.Printf("unsubscribed pull/%s", topic)

	// Destroy after last topic
	if len(c.List) == 0 {
		err = destroyPullPoint(c.PullPt)
		c.PullPt = ""
	}
	return
}

// Start set the active flag on all topics
func (c *PullClient) Start() (err error) {
	for _, topic := range c.List {
		topic.Started = true
	}
	return
}

// Stop unsubscribe from all topics
func (c *PullClient) Stop() (err error) {
	for name, topic := range c.List {
		if topic.Started {
			err = c.Unsubscribe(name)
		}
	}
	return
}

// Read implements the io.Reader interface
func (c *PullClient) Read(p []byte) (n int, err error) {
	data := "pull/foobar"
	n = len(data)
	copy(p, []byte(data))
	err = io.EOF
	return
}
