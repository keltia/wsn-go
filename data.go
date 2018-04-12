// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import "encoding/xml"

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

// io.Reader interface
func (d *WsnData) Read(p []byte) (n int, err error) {
	n = copy(p, d.Body.Notify.NotificationMessage.Message.Data)
	err = nil
	return
}

func (d *WsnData) String() string {
	return string(d.Body.Notify.NotificationMessage.Message.Data)
}
