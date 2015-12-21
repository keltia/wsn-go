// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
    "testing"
    "encoding/xml"
	"strings"
)

const (
    testRecord = `
<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"
               xmlns:wsa="http://www.w3.org/2005/08/addressing">
   <soap:Body>
      <b:Notify xmlns:b="http://docs.oasis-open.org/wsn/b-2">
         <b:NotificationMessage>
            <b:SubscriptionReference>
               <wsa:Address
                > http://swan.eurocontrol.fr:9000/wsn/subscriptions/ID-147-196-221-30-149aec1e243-0-0</wsa:Address>
               <wsa:ReferenceParameters/>
               <wsa:Metadata>
                  <wsam:ServiceName EndpointName="PausableSubscriptionManagerPort"
                                    xmlns:wsam="http://www.w3.org/2007/05/addressing/metadata"
                                    xmlns:ns2="http://docs.oasis-open.org/wsn/bw-2"
                   >ns2:PausableSubscriptionManager</wsam:ServiceName>
               </wsa:Metadata>
            </b:SubscriptionReference>
            <b:Topic Dialect="http://docs.oasis-open.org/wsn/t-1/TopicExpression/Simple"
             >Unaligned-PER</b:Topic>
            <b:Message>
               <adsc-uper:PDU xmlns:adsc-uper="http://www.eurocae.net/ADS_C_vH_stepC/Unaligned-PER">
                  <type xmlns="">ADSReport</type>
                  <value xmlns="">8009A600C426A0901772359A60100018</value>
               </adsc-uper:PDU>
            </b:Message>
         </b:NotificationMessage>
      </b:Notify>
   </soap:Body>
</soap:Envelope>
`

   testResult = `<adsc-uper:PDU xmlns:adsc-uper="http://www.eurocae.net/ADS_C_vH_stepC/Unaligned-PER">
                  <type xmlns="">ADSReport</type>
                  <value xmlns="">8009A600C426A0901772359A60100018</value>
               </adsc-uper:PDU>
`
)

// I wrote mine to get the exact byte difference
func diffStrings(str1, str2 string) (off int, a byte, b byte) {
	bstr1 := []byte(str1)
	bstr2 := []byte(str1)
	if len(bstr1) == len(bstr2) {
		for i, _ := range bstr1 {
			if bstr1[i] != bstr2[i] {
				off, a, b = i, byte(bstr1[i]), byte(bstr2[i])
				return
			}
		}
	} else {
		off, a, b = -2, byte(42), byte(42)
	}
	return -1, byte(42), byte(42)
}

func TestWsnStringer(t *testing.T) {
    var wsnTest WsnData

    err := xml.Unmarshal([]byte(testRecord), &wsnTest)
    if err != nil {
        t.Errorf("Data can not be decoded: %v\nError: %v", testRecord, err)
    }

    str1 := wsnTest.String()
	str2 := strings.TrimSpace(testResult)
	diff, a, b := diffStrings(str1, str2)
    if diff == -2 {
		t.Errorf("Error: lengths are different: %d %d", len(str1), len(str2))
	} else if diff != -1 {
        t.Errorf("Decoded data is different:\nByte at %d (%d) != Byte at %d (%d)", diff, a, diff, b)
    }
}
