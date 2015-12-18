package wsn

import (
	"encoding/xml"
	"github.com/keltia/wsn-go/config"
)

// My stuff

type SubVars struct {
	TopicName string
	TopicURL  string
}

type Client struct {
	Config   *config.Config
	Target   string
	Topics   map[string]*Topic
	Feed_one func([]byte)
	Verbose  bool
	Timeout  int64
}

type Topic struct {
	Bytes     int64
	Pkts      int
	UnsubAddr string
	Started   bool
}

// SOAP stuff

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

// Generic WS-N data

type WsnData struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Notify struct {
			NotificationMessage struct {
				SubscriptionReference SAReference
				Topic           string
				Message         struct {
					Data []byte `xml:",innerxml"`
				}
			}
		}
	}
}

func (d *WsnData) String() string {
	return string(d.Body.Notify.NotificationMessage.Message.Data)
}