// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package soap

import (
    "testing"
	"strings"
)

var testVars = SubVars{
	TopicURL: "http://example.com/foo",
	TopicName: "foobar",
}

const (
	testSubscribePush = `
<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
	       xmlns:b="http://docs.oasis-open.org/wsn/b-2"
	       xmlns:add="http://www.w3.org/2005/08/addressing">
   <soap:Header/>
   <soap:Body>
      <b:Subscribe>
	 <b:ConsumerReference>
	    <add:Address>http://example.com/foo</add:Address>
	 </b:ConsumerReference>
	 <b:Filter>
	   <b:TopicExpression>foobar</b:TopicExpression>
	 </b:Filter>
      </b:Subscribe>
   </soap:Body>
</soap:Envelope>
`

    emptySubscribePush = `
<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
	       xmlns:b="http://docs.oasis-open.org/wsn/b-2"
	       xmlns:add="http://www.w3.org/2005/08/addressing">
   <soap:Header/>
   <soap:Body>
      <b:Subscribe>
	 <b:ConsumerReference>
	    <add:Address></add:Address>
	 </b:ConsumerReference>
	 <b:Filter>
	   <b:TopicExpression></b:TopicExpression>
	 </b:Filter>
      </b:Subscribe>
   </soap:Body>
</soap:Envelope>
`
)

func TestNewRequest(t *testing.T) {
	req, err := NewRequest(SUBSCRIBEPUSH, testVars)
	if err != nil {
		t.Errorf("Templ creation should not fail: %v", err)
	}

	breq := strings.TrimSpace(req.Text.String())
	btest := strings.TrimSpace(testSubscribePush)
	if breq != btest {
		t.Errorf("Error: different output:\n%v\n--\n%v", breq, btest)
	}
	if req.Action != SUBSCRIBEPUSH {
		t.Errorf("Error: bad action: %v", req.Action)
	}

	req, err = NewRequest(SUBSCRIBEPUSH, SubVars{})
	if err != nil {
		t.Errorf("Templ creation should not fail w/empty subs")
	}
	breq = strings.TrimSpace(req.Text.String())
	btest = strings.TrimSpace(emptySubscribePush)
	if breq != btest {
		t.Errorf("Error: different output:\n%v\n--\n%v", breq, btest)
	}
	if req.Action != SUBSCRIBEPUSH {
		t.Errorf("Error: bad action: %v", req.Action)
	}


	req, err = NewRequest(UNSUBSCRIBEPUSH, testVars)

	breq = strings.TrimSpace(req.Text.String())
	btest = strings.TrimSpace(unsubscribePushText)
	if breq != btest {
		t.Errorf("Error: different output:\n%v\n--\n%v", breq, btest)
	}
	if req.Action != UNSUBSCRIBEPUSH {
		t.Errorf("Error: bad action: %v", req.Action)
	}

}