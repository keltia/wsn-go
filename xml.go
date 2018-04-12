// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"encoding/xml"
)

const (

)

// SOAP stuff

type SubVars struct {
	TopicName string
	TopicURL  string
}

type SubscribeAnswer struct {
	XMLName xml.Name
	Body    SABody
}

type SABody struct {
	XMLName xml.Name
	Resp    SAResp `xml:"SubscribeResponse"`
}

type SAResp struct {
	XMLName   xml.Name    `xml:"SubscribeResponse"`
	Reference SAReference `xml:"SubscriptionReference"`
}

type SAReference struct {
	XMLName             xml.Name `xml:"SubscriptionReference"`
	Address             string
	ReferenceParameters string
	Metadata            string
}

