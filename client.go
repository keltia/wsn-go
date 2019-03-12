// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

const (
	MAX_TOPICS = 10

	MODE_PULL = 1 + iota
	MODE_PUSH
)

// Generic Client interface for both Push/Pull modes
type Client interface {
	Type() int
	Subscribe(string) error
	Unsubscribe(string) error
	Start() error
	Stop() error
}

