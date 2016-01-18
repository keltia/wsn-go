// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package soap

import (
	"errors"
)

// Unknown operation
var ErrUnknownOperation = errors.New("Unknown operation")

// Error when creating template
var ErrCantCreateTemplate = errors.New("Can not create template")

// No template for given action
var ErrTemplateNotFound = errors.New("Unknown action")
