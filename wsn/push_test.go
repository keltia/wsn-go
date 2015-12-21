// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"testing"
	"fmt"
	"wsn-go/config"
)

var _topics TopicList = TopicList{}

func TestNewPushClient(t *testing.T) {
	// Load our stuff
	config, err :=  config.LoadConfig("../config/config.toml")
	if err != nil {
		t.Errorf("Error loading config: %v", err)
	}
	client := NewPushClient(config)

	// Check various fields

	// Public
	t.Log("  Public fields")
	if len(client.List) != 0 {
		t.Errorf("List is not correctly set: %s", client.List)
	}
	if client.Timeout != int64(-1) {
		t.Errorf("Timeout is not correctly set: %s", client.Timeout)
	}

	// Private
	t.Log("  Private fields")
	if client.base != config.Base {
		t.Errorf("base is not correctly set: %s", client.base)
	}
	if client.target != fmt.Sprintf("%s:%d/%s", config.Target, config.Port, config.Broker) {
		t.Errorf("target is not correctly set: %s", client.target)
	}
	if client.port != config.Port {
		t.Errorf("port is not correctly set: %s", client.port)
	}
	if client.verbose != false {
		t.Errorf("verbose is not correctly set: %s", client.verbose)
	}
}

func TestPushType (t *testing.T) {
	// Load our stuff
	config, err :=  config.LoadConfig("../config/config.toml")
	if err != nil {
		t.Errorf("Error loading config: %v", err)
	}
	client := NewPushClient(config)

	if client.Type() != MODE_PUSH {
		t.Errorf("Wrong type for %v", client)
	}
}
