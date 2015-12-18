// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"log"
	"io"
)

// A PullClient represents an active Pull mode client for WS-N.  It maintains a list of
// subscribed topics and the Pull point that will be used to subscribe.
type PullClient struct {
	PullPt string
	List   TopicList
}

// createPullPoint create a Pull point on the broker that will be used to subscribe
// topics.
func createPullPoint() (pullPt string, err error) {
	return
}

// destroyPullPoint de-registers the pull point to avoid hogging resources on the broker
func destroyPullPoint(pullPt string) (err error) {
	return
}

// NewPullClient creates a new instance of a Pull client.
func NewPullClient() *PullClient {
	return &PullClient{
		PullPt: "",
		List:   TopicList{},
	}
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
		Started: false,
		UnsubAddr: "",
		Bytes: 0,
		Pkts: 0,
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

