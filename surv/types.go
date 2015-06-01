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

// Surv data (CAT62 simplified)

type SurvData struct {
    XMLName xml.Name
    Body    SDBody
}

type SDBody struct {
    XMLName     xml.Name
    Notify      SDNotify `xml:"Notify"`
}

type SDNotify struct {
	XMLName		xml.Name
	Message		SDMessage `xml:"NotificationMessage"`
}

// Data can be JSON, compressed JSON or XML

type SDMessage struct {
	XMLName		xml.Name
	Topic		string `xml:"SubscriptionReference"`
	Payload		SDPayload
}

type SDPayload struct {
	XMLName		xml.Name
	Data		[]byte `xml:"PlainText"`
}
