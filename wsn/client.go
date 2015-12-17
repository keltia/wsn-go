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

type Client interface {
	io.Reader

	Type() int
	Subscribe(string) error
	Unsubscribe(string) error
	Start() error
	Stop() error
}

// Error list

var ErrTopicNotFound = errors.New("Topic not found")

var ErrCreatingPullPoint = errors.New("Can not create pull point")
