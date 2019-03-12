// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"encoding/xml"
)

const ()

// SOAP stuff

type SubscribeAnswer struct {
	XMLName xml.Name
	Body    []byte `xml:", innerxml"`
}

// PushClient.Subscribe
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

// PullClient.createPullPoint
type CPPBody struct {
	XMLName xml.Name
	CreatePullPointResponse struct{
		PullPt struct {
			Address string
			Params string
			Metadata []byte `xml:",innerxml"`
		}
	}
}

// PullClient.realSubscribe
type STPPBody struct {
	XMLName xml.Name
	SubscribeTopicResponse struct {
		Reference struct {
			Address string
			Params string
			Metadata []byte `xml:",innerxml"`
		}
    }
}
