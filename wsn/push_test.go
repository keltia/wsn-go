// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"fmt"
	"testing"
	"wsn-go/config"
)

var _topics TopicList = TopicList{}

func TestNewPushClient(t *testing.T) {
	// Load our stuff
	config, err := config.LoadConfig("../config/config.toml")
	if err != nil {
		t.Errorf("Error loading config: %v", err)
	}
	client := NewPushClient(config)

	// Check various fields

	// Public
	if len(client.List) != 0 {
		t.Errorf("List is not correctly set: %v", client.List)
	}
	if client.Timeout != int64(-1) {
		t.Errorf("Timeout is not correctly set: %d", client.Timeout)
	}

	// Private
	if client.base != config.Base {
		t.Errorf("base is not correctly set: %s", client.base)
	}
	if client.target != fmt.Sprintf("%s:%d/%s", config.Target, config.Port, config.Broker) {
		t.Errorf("target is not correctly set: %s", client.target)
	}
	if client.port != config.Port {
		t.Errorf("port is not correctly set: %d", client.port)
	}
	if client.verbose != false {
		t.Errorf("verbose is not correctly set: %v", client.verbose)
	}
}

func TestPushType(t *testing.T) {
	// Load our stuff
	config, err := config.LoadConfig("../config/config.toml")
	if err != nil {
		t.Errorf("Error loading config: %v", err)
	}
	client := NewPushClient(config)

	if client.Type() != MODE_PUSH {
		t.Errorf("Wrong type for %v", client)
	}
}

func TestPushSetVerbose(t *testing.T) {
	// Load our stuff
	config, err := config.LoadConfig("../config/config.toml")
	if err != nil {
		t.Errorf("Error loading config: %v", err)
	}
	client := NewPushClient(config)

	client.SetVerbose()
	if !client.verbose {
		t.Errorf("Error setting verbose mode: %v", client.verbose)
	}
}

func TestPushSubscribe(t *testing.T) {
	// Load our stuff
	config, err := config.LoadConfig("../config/config.toml")
	if err != nil {
		t.Errorf("Error loading config: %v", err)
	}
	client := NewPushClient(config)

	err = client.Subscribe("foobar")
	if err != nil {
		t.Errorf("Error: Subscribe returned error: %v", err)
	}

	// do we have exactly one topic?
	if len(client.List) != 1 {
		t.Errorf("Error: List should have 1 item: %d", len(client.List))
	}

	// is ours present?
	if _, ok := client.List["foobar"]; !ok {
		t.Errorf("Error: topic foobar not present: %v", client.List)
	}

	// does it have the right settings?
	topic := client.List["foobar"]
	if topic.Started {
		t.Errorf("Error: topic should NOT be started")
	}
	if topic.UnsubAddr != "" {
		t.Errorf("Error: unsubaddr should be empty!: %s", topic.UnsubAddr)
	}
}

func TestPushUnsubscribe(t *testing.T) {
	// Load our stuff
	config, err := config.LoadConfig("../config/config.toml")
	if err != nil {
		t.Errorf("Error loading config: %v", err)
	}
	client := NewPushClient(config)

	err = client.Subscribe("foobar")
	if err != nil {
		t.Errorf("Error: Subscribe returned error: %v", err)
	}

	//err = client.Unsubscribe("foobar")
}

func TestSetTimeout(t *testing.T) {
	// Load our stuff
	config, err := config.LoadConfig("../config/config.toml")
	if err != nil {
		t.Errorf("Error loading config: %v", err)
	}
	client := NewPushClient(config)

	client.SetTimeout(255)
	if client.Timeout != 255 {
		t.Errorf("Error setting timeout: %d - %d", 255, client.Timeout)
	}
}
