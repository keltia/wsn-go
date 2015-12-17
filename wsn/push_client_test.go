package wsn

import (
	"testing"
	"github.com/keltia/wsn-go/config"
)

var pushConfig = config.Config{Proto: "http", Site: "example.com", Port: 666, Endpoint: "foo"}

var push_topics map[string]*Topic
var push_topic *Topic

func TestNewPushClient(t *testing.T) {
	var err error

	client, err := NewPushClient(&pushConfig)
	if err != nil {
		t.Errorf("Bad init: %v: %v", client, err)
	}

	conf := client.Config
	if conf.Port != pushConfig.Port || conf.Site != pushConfig.Site {
		t.Errorf("Wrong fields: %v: %v", conf, client)
	}

	if client.Verbose != false {
		t.Errorf("Error: Verbose should be false! %v", client.Verbose)
	}

	if conf.Proto != pushConfig.Proto || conf.Endpoint != pushConfig.Endpoint {
		t.Errorf("Wrong fields: %v: %v", conf, client)
	}

	if client.Target != generateEndpoint(conf) {
		t.Errorf("Wrong Target: %v: %v", conf, client)
	}

	if client.Topics == nil {
		t.Errorf("Uninitialized topics: %v: %v", client)
	}
}

func TestAddFeed(t *testing.T) {
	client, err := NewPushClient(&pushConfig)
	if err != nil {
		t.Errorf("Bad init: %v: %v", client, err)
	}

	client.AddFeed("foobar")
	topic := client.Topics["foobar"]

	if topic == nil {
		t.Errorf("Error: client.Topics[\"foobar\"] is null")
	}

	if topic.Started != false {
		t.Errorf("Error: topic should not be startted")
	}
}

func TestSetTimer(t *testing.T) {
	client, err := NewPushClient(&pushConfig)
	if err != nil {
		t.Errorf("Bad init: %v: %v", client, err)
	}

	tval := int64(3600)
	client.SetTimer(tval)
	if client.Timeout != tval {
		t.Errorf("Bad init for Timeout: %v: %v", client, client.Timeout)
	}
}

func TestAddHandler(t *testing.T) {
	client, err := NewPushClient(&pushConfig)
	if err != nil {
		t.Errorf("Bad init: %v: %v", client, err)
	}

	client.AddHandler(nil)
	if client.Feed_one != nil {
		t.Errorf("Error: error setting Feed_one()\n")
	}
}
