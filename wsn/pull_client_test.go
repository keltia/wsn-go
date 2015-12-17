package wsn

import (
	"testing"
	"github.com/keltia/wsn-go/config"
)

var pullConfig = config.Config{Proto: "http", Site: "example.com", Port: 666, Endpoint: "foo"}

var pull_topics map[string]*Topic
var pull_topic *Topic

func TestNewPullClient(t *testing.T) {
	var err error

	client, err := NewPullClient(&pullConfig)
	if err != nil {
		t.Errorf("Bad init: %v: %v", client, err)
	}

	conf := client.Config
	if conf.Port != pullConfig.Port || conf.Site != pullConfig.Site {
		t.Errorf("Wrong fields: %v: %v", conf, client)
	}

	if client.Verbose != false {
		t.Errorf("Error: Verbose should be false! %v", client.Verbose)
	}

	if conf.Proto != pullConfig.Proto || conf.Endpoint != pullConfig.Endpoint {
		t.Errorf("Wrong fields: %v: %v", conf, client)
	}

	if client.Target != generateEndpoint(conf) {
		t.Errorf("Wrong Target: %v: %v", conf, client)
	}

	if client.Topics == nil {
		t.Errorf("Uninitialized topics: %v: %v", client)
	}
}

