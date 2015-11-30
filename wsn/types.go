package wsn

import (
	"../config"
	"encoding/xml"
)

// My stuff

type SubVars struct {
	TopicName	string
	TopicURL	string
}

type Client struct {
	Config		*config.Config
	Target		string
	Wsdl		string
	Topics		map[string]Topic
	Feed_one	func([]byte)
	Verbose		bool
}

type Topic struct {
	Bytes		int64
	Pkts		int
	UnsubAddr	string
	Started		bool
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

// Surv data (CAT62 simplified)

type SurvData struct {
    XMLName xml.Name `xml:"Envelope"`
    Body    SDBody
}

type SDBody struct {
    Notify      SDNotify
}

type SDNotify struct {
	NotifyMsg	SDNotifyMsg `xml:"NotificationMessage"`
}

// Data can be JSON, compressed JSON or XML

type SDNotifyMsg struct {
	NotifyRef	SAReference `xml:"SubscriptionReference"`
	Topic		string
	Message		[]byte		// Generic payload
}
