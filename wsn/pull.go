package wsn

import (
	"log"
	"io"
)

type PullClient struct {
	PullPt string
	List   TopicList
}

func createPullPoint() (pullPt string, err error) {
	return
}

func destroyPullPoint(pullPt string) (err error) {
	return
}

func NewPullClient() *PullClient {
	return &PullClient{
		PullPt: "",
		List:   TopicList{},
	}
}

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

func (c *PullClient) Unsubscribe(topic string) (err error) {
	log.Printf("unsubscribed pull/%s", topic)

	// Destroy after last topic
	if len(c.List) == 0 {
		err = destroyPullPoint(c.PullPt)
		c.PullPt = ""
	}
	return
}

func (c *PullClient) Start() (err error) {
	return
}

func (c *PullClient) Stop() (err error) {
	for name, _ := range c.List {
		err = c.Unsubscribe(name)
	}
	return
}

func (c *PullClient) Read(p []byte) (n int, err error) {
	data := "pull/foobar"
	n = len(data)
	copy(p, []byte(data))
	err = io.EOF
	return
}

