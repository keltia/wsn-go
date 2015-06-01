package surv

import (
	"../config"
	"encoding/xml"
)

// My stuff

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

// SOAP stuff

type SubscribeAnswer struct {
    XMLName xml.Name
    Body    SABody
}

type SABody struct {
    XMLName     xml.Name
    Resp        SAResp `xml:"SubscribeResponse"`
}

type SAResp struct {
	XMLName	xml.Name `xml:"SubscribeResponse"`
	Reference	SAReference `xml:"SubscriptionReference"`
}

type SAReference struct {
	XMLName xml.Name `xml:"SubscriptionReference"`
	Address	string
	ReferenceParameters string
	Metadata string
}

