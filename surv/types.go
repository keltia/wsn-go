package surv

import "../config"

type SubVars struct {
	my_topic	string
	topic		string
}

type Client struct {
	config		config.Config
	Target		string
	Wsdl		string
	Topics		map[string]Topic
	Feed_one	func([]byte)
}

type Topic struct {
	Bytes	int64
	Pkts	int
	Address	string
	Started	bool
}

