// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"fmt"
	"testing"

	"wsn-go/config"
)

func TestNewPullClient(t *testing.T) {
	// Load our stuff
	config, err := config.LoadConfig("../config/config.toml")
	if err != nil {
		t.Errorf("Error loading config: %v", err)
	}
	client := NewPullClient(config)

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

func TestPullType(t *testing.T) {
	// Load our stuff
	config, err := config.LoadConfig("../config/config.toml")
	if err != nil {
		t.Errorf("Error loading config: %v", err)
	}
	client := NewPullClient(config)

	if client.Type() != MODE_PULL {
		t.Errorf("Wrong type for %v", client)
	}
}

func TestPullSetVerbose(t *testing.T) {
	// Load our stuff
	config, err := config.LoadConfig("../config/config.toml")
	if err != nil {
		t.Errorf("Error loading config: %v", err)
	}
	client := NewPullClient(config)

	client.SetVerbose()
	if !client.verbose {
		t.Errorf("Error setting verbose mode: %v", client.verbose)
	}
}

