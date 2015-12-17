package wsn

import (
	"testing"
	"github.com/keltia/wsn-go/config"
)

var myConfig = config.Config{Proto: "http", Site: "example.com", Port: 666, Endpoint: "foo"}
//var client wsn.Client

var _topics map[string]*Topic
var _topic *Topic

func TestNewPushClient(t *testing.T) {
	var err error

	client, err := NewPushClient(&myConfig)
	if err != nil {
		t.Errorf("Bad init: %v: %v", client, err)
	}

	conf := client.Config
	if conf.Port != myConfig.Port || conf.Site != myConfig.Site {
		t.Errorf("Wrong fields: %v: %v", conf, client)
	}

	if client.Verbose != false {
		t.Errorf("Error: Verbose should be false! %v", client.Verbose)
	}

	if conf.Proto != myConfig.Proto || conf.Endpoint != myConfig.Endpoint {
		t.Errorf("Wrong fields: %v: %v", conf, client)
	}

	if client.Target != client.generateURL(conf.Endpoint) {
		t.Errorf("Wrong Target: %v: %v", conf, client)
	}

	if client.Topics == nil {
		t.Errorf("Uninitialized topics: %v: %v", client)
	}
}

func TestAddFeed(t *testing.T) {
	client, err := NewPushClient(&myConfig)
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
	client, err := NewPushClient(&myConfig)
	if err != nil {
		t.Errorf("Bad init: %v: %v", client, err)
	}

	tval := int64(3600)
	client.SetTimer(tval)
	if client.Timeout != tval {
		t.Errorf("Bad init for Timeout: %v: %v", client, client.Timeout)
	}
}

func TestGenerateURL(t *testing.T) {
	client, err := NewPushClient(&myConfig)
	if err != nil {
		t.Errorf("Bad init: %v: %v", client, err)
	}

	url := client.generateURL(myConfig.Endpoint)
	if url != "http://example.com:666/foo" {
		t.Errorf("Error: bad format %s for %v\n", url, client)
	}
}

func TestAddHandler(t *testing.T) {
	client, err := NewPushClient(&myConfig)
	if err != nil {
		t.Errorf("Bad init: %v: %v", client, err)
	}

	client.AddHandler(nil)
	if client.Feed_one != nil {
		t.Errorf("Error: error setting Feed_one()\n")
	}
}
