// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"encoding/xml"
	"strings"
)

// Generic WS-N data

type WsnData struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Notify struct {
			NotificationMessage struct {
				SubscriptionReference SAReference
				Topic                 string
				Message               struct {
					Data []byte `xml:",innerxml"`
				}
			}
		}
	}
}

func (d *WsnData) String() (str string) {
	str = strings.TrimSpace(string(d.Body.Notify.NotificationMessage.Message.Data))
	return
}
