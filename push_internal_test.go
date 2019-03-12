// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
)

var (
	pushSADocument = `<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
	               xmlns:wsa="http://www.w3.org/2005/08/addressing">
	   <soap:Body>
	      <CreatePullPointResponse xmlns="http://docs.oasis-open.org/wsn/b-2">
	         <PullPoint>
	            <wsa:Address>http://swan.eurocontrol.fr:9000/wsn/pullpoints/ID14719622130149aec1e24310</wsa:Address>
	            <wsa:ReferenceParameters/>
	            <wsa:Metadata>
	               <wsam:ServiceName EndpointName="PullPointPort"
	                           xmlns:wsam="http://www.w3.org/2007/05/addressing/metadata"
	                           xmlns:ns2="http://docs.oasis-open.org/wsn/bw-2">ns2:PullPoint</wsam:ServiceName>
	            </wsa:Metadata>
	         </PullPoint>
	      </CreatePullPointResponse>
	   </soap:Body>
	</soap:Envelope>
	`
)

