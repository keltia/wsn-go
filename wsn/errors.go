// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"errors"
)

// Error list

// Channel errors
var ErrNoOutputChannel = errors.New("Output channel not opened")

// Topic errors
var ErrTopicNotFound = errors.New("Topic not found")

// Can't add an existing topic
var ErrTopicAlreadyExist = errors.New("Topic already exist!")

// Can't subscribe more than once
var ErrTopicAlreadySubscribed = errors.New("Topic already subscribed!")

// Pull point creation error
var ErrCreatingPullPoint = errors.New("Can not create pull point")

// Pull Point already created
var ErrPullPtAlreadyPresent = errors.New("Pull point already present")

// Pull point destruction errors
var ErrDestroyingPullPoint = errors.New("Can not destroy pull point")

// Templating error
var ErrCantCreateTemplate = errors.New("Can not instanciate template")
