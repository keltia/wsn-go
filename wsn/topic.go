package wsn

type Topic struct {
	Started   bool
	UnsubAddr string
	Bytes     int64
	Pkts      int
}

type TopicList map[string]*Topic
