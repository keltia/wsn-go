// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"errors"
)

// Error list

// ErrNoOutputChannel is for Channel errors
var ErrNoOutputChannel = errors.New("output channel not opened")

// ErrTopicNotFound is for Topic errors
var ErrTopicNotFound = errors.New("topic not found")

// ErrTopicAlreadyExist is when we can't add an existing topic
var ErrTopicAlreadyExist = errors.New("topic already exist")

// ErrTopicAlreadySubscribed is when we can't subscribe more than once
var ErrTopicAlreadySubscribed = errors.New("topic already subscribed")

// ErrCreatingPullPoint is when there is a Pull point creation error
var ErrCreatingPullPoint = errors.New("can not create pull point")

// ErrPullPtAlreadyPresent is when Pull Point is already created
var ErrPullPtAlreadyPresent = errors.New("pull point already present")

// ErrDestroyingPullPoint is for Pull point destruction errors
var ErrDestroyingPullPoint = errors.New("can not destroy pull point")

// ErrCantSubscribeTopicPull is for subscription errors
var ErrCantSubscribeTopicPull = errors.New("can not subscribe pullpt to topic")

// ErrCantCreateTemplate is for Templating errors
var ErrCantCreateTemplate = errors.New("can not instanciate template")
