// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"bytes"
	"text/template"
)

const (
	// Subscribe to topic
	// sent to: wsn/NotificationBroker
	// answer: wsn/subscriptions/<id>
	subscribePushText = `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
	               xmlns:b="http://docs.oasis-open.org/wsn/b-2"
	               xmlns:add="http://www.w3.org/2005/08/addressing">
	   <soap:Header/>
	   <soap:Body>
	      <b:Subscribe>
	         <b:ConsumerReference>
	            <add:Address>{{.TopicURL}}</add:Address>
	         </b:ConsumerReference>
	         <b:Filter>
	           <b:TopicExpression>
	             {{.TopicName}}
	           </b:TopicExpression>
	         </b:Filter>
	      </b:Subscribe>
	   </soap:Body>
	</soap:Envelope>
	`

	// Unsubscribe from topic
	// sent to: wsn/subscriptions/<id>
	unsubscribePushText = `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
	               xmlns:b="http://docs.oasis-open.org/wsn/b-2">
	   <soap:Header/>
	   <soap:Body>
	      <b:Unsubscribe/>
	   </soap:Body>
	</soap:Envelope>
	`

	// Create Pull Point
    // sent to: wsn/CreatePullPoint
    // answer: wsn/pullpoints/<id>
	createPullPointText = `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
	               xmlns:wsa="http://www.w3.org/2005/08/addressing">
	   <soap:Header>
	      <wsa:Action>
	        http://docs.oasis-open.org/wsn/bw-2/CreatePullPoint/CreatePullPointRequest
	      </wsa:Action>
	   </soap:Header>
	   <soap:Body>
	      <b:CreatePullPoint xmlns:b="http://docs.oasis-open.org/wsn/b-2"/>
	   </soap:Body>
	</soap:Envelope>
	`

	// Subscribe topic to Pull Point
    // sent to: wsn/NotificationBroker
	// answer: wsn/subscriptions/<id>
	subscribePullPointText = `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
	               xmlns:wsa="http://www.w3.org/2005/08/addressing">
	   <soap:Header>
	      <wsa:Action>
	        http://docs.oasis-open.org/wsn/bw-2/NotificationProducer/SubscribeRequest
	      </wsa:Action>
	   </soap:Header>
	   <soap:Body>
	      <b:Subscribe xmlns:b="http://docs.oasis-open.org/wsn/b-2">
	         <b:ConsumerReference>
	            <wsa:Address>
	              {{.PullPt}}
	            </wsa:Address>
	         </b:ConsumerReference>
	         <b:Filter>
	            <b:TopicExpression Dialect="http://docs.oasis-open.org/wsn/t-1/TopicExpression/Simple">
	              {{.Topic}
	            </b:TopicExpression>
	         </b:Filter>
	      </b:Subscribe>
	   </soap:Body>
	</soap:Envelope>
	`

	// Get messages from topics
    // sent to: wsn/pullpoints/<id>
	getMessagePullText = `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
	               xmlns:wsa="http://www.w3.org/2005/08/addressing">
	   <soap:Header>
	      <wsa:Action>http://docs.oasis-open.org/wsn/bw-2/PullPoint/GetMessagesRequest</wsa:Action>
	   </soap:Header>
	   <soap:Body>
	      <b:GetMessages xmlns:b="http://docs.oasis-open.org/wsn/b-2">
	         <b:MaximumNumber>{{.HowMany}}</b:MaximumNumber>
	      </b:GetMessages>
	   </soap:Body>
	</soap:Envelope>
	`

	// Unsubscribe topic
    // sent to: wsn/subscriptions/<id>
	unsubscribePullPointText = `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
	               xmlns:wsa="http://www.w3.org/2005/08/addressing">
	   <soap:Header>
	      <wsa:Action>
	        http://docs.oasis-open.org/wsn/bw-2/SubscriptionManager/UnsubscribeRequest
	      </wsa:Action>
	   </soap:Header>
	   <soap:Body>
	      <b:Unsubscribe xmlns:b="http://docs.oasis-open.org/wsn/b-2"/>
	   </soap:Body>
	</soap:Envelope>
	`

	// Destroy Pull Point
    // sent to: wsn/pullpoints/<id>
	destroyPullPointText = `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
	               xmlns:wsa="http://www.w3.org/2005/08/addressing">
	   <soap:Header>
	     <wsa:Action>http://docs.oasis-open.org/wsn/bw-2/PullPoint/DestroyPullPointRequest</wsa:Action>
	   </soap:Header>
	   <soap:Body>
	     <b:DestroyPullPoint xmlns:b="http://docs.oasis-open.org/wsn/b-2"/>
	   </soap:Body>
	</soap:Envelope>
	`
)

func createRequest(action, templ string, vars SubVars) (result bytes.Buffer, err error) {
	t := template.Must(template.New(action).Parse(templ))
	if err = t.Execute(&result, vars); err != nil {
		err = ErrCantCreateTemplate
		result = bytes.Buffer{}
	}
	return
}
