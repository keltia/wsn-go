// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

type Topic struct {
	Started   bool
	UnsubAddr string
	Bytes     int64
	Pkts      int
}

type TopicList map[string]*Topic
