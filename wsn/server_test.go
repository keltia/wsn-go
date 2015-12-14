package wsn

import (
	"encoding/xml"
	"testing"
	//"github.com/keltia/wsn-go/config"
)

const (
	onePkt = `
<?xml version="1.0"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
        <soap:Body>
                <ns2:Notify xmlns="http://www.w3.org/2005/08/addressing" xmlns:ns2="http://docs.oasis-open.org/wsn/b-2" xmlns:ns3="http://docs.oasis-open.org/wsrf/bf-2" xmlns:ns4="http://docs.oasis-open.org/wsrf/rp-2" xmlns:ns5="http://docs.oasis-open.org/wsn/t-1" xmlns:ns6="http://docs.oasis-open.org/wsn/br-2" xmlns:ns7="http://docs.oasis-open.org/wsrf/r-2">
                        <ns2:NotificationMessage>
                                <ns2:SubscriptionReference>
                                        <Address>http://0.0.0.0:9000/wsn/subscriptions/ID-192-70-89-113-14c32c59067-0-1</Address>
                                        <ReferenceParameters/>
                                        <Metadata>
                                                <wsam:ServiceName xmlns:wsam="http://www.w3.org/2007/05/addressing/metadata" xmlns:ns2="http://services.wsn.cxf.apache.org/" xmlns:wsa="http://www.w3.org/2005/08/addressing" xmlns:wsaw="http://www.w3.org/2006/05/addressing/wsdl" EndpointName="JaxwsSubscriptionPort">ns2:JaxwsSubscriptionService</wsam:ServiceName>
                                        </Metadata>
                                </ns2:SubscriptionReference>
                                <ns2:Topic>
                   AsterixJSON
                 </ns2:Topic>
                                <ns2:Message>
                                        <surv:Cat62SurveillanceJSON xmlns:surv="http://www.eurocontrol.int/sesar/swim-mc/surv" xmlns:wsnt="http://docs.oasis-open.org/wsn/b-2">
                                                <surv:PlainText>{"Cat62Surveillance":{"SurveillanceRecord":[]}}</surv:PlainText>
                                        </surv:Cat62SurveillanceJSON>
                                </ns2:Message>
                        </ns2:NotificationMessage>
                </ns2:Notify>
        </soap:Body>
</soap:Envelope>
	`
	decodedPkt = `<surv:Cat62SurveillanceJSON xmlns:surv="http://www.eurocontrol.int/sesar/swim-mc/surv" xmlns:wsnt="http://docs.oasis-open.org/wsn/b-2">
<surv:PlainText>{"Cat62Surveillance":{"SurveillanceRecord":[]}}</surv:PlainText></surv:Cat62SurveillanceJSON>
`
)

func TestDecodeNotify(t *testing.T) {
	body := []byte(onePkt)
	notify := &WsnData{}
	err := xml.Unmarshal(body, notify)
	if err != nil {
		t.Errorf("Payload should be identical: %v\n%v", decodedPkt, notify)
	}
}
