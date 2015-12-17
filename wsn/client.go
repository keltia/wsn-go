package wsn

import (
	"io"
	"errors"
)

const (
	MAX_TOPICS = 10
)

type Client interface {
	io.Reader

	Subscribe(string) error
	Unsubscribe(string) error
	Start() error
	Stop() error
}

// Error list

var ErrTopicNotFound = errors.New("Topic not found")

var ErrCreatingPullPoint = errors.New("Can not create pull point")
