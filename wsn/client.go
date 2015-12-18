// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"io"
	"errors"
)

const (
	MAX_TOPICS = 10

	MODE_PULL = 1 + iota
	MODE_PUSH
)

// Generic Client interface for both Push/Pull modes
type Client interface {
	io.Reader

	Type() int
	Subscribe(string) error
	Unsubscribe(string) error
	Start() error
	Stop() error
}

// Error list

// Topic errors
var ErrTopicNotFound = errors.New("Topic not found")

// Pull point creation error
var ErrCreatingPullPoint = errors.New("Can not create pull point")

// Pull point destruction errors
var ErrDestroyingPullPoint = errors.New("Can not destroy pull point")
