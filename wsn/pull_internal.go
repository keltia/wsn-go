// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"encoding/xml"
	"github.com/keltia/wsn-go/soap"
)

// createPullPoint create a Pull point on the broker that will be used to subscribe
// topics.
func (c *PullClient) createPullPoint() (pullPt string, err error) {

	// Pull point should not exist at this stage
	if c.PullPt != "" {
		soapReq, err := soap.NewRequest(soap.CREATEPULLPOINT, soap.SubVars{})
		if err != nil {
			return "", ErrCreatingPullPoint
		}
		answer, err := soapReq.Send(c.target)
		if err != nil {
			var res *CPPBody

			err = xml.Unmarshal(answer, res)
		}
	} else {
		return "", ErrPullPtAlreadyPresent
	}
	return c.PullPt, nil
}

// destroyPullPoint de-registers the pull point to avoid hogging resources on the broker
func (c *PullClient) destroyPullPoint(pullPt string) (err error) {
	if false {
		err = ErrDestroyingPullPoint
	}
	return
}

// realSubscribe does the registration for the given topic
func (c *PullClient) realSubscribe(topic string) (err error) {

	// Prepare the request
	var soapReq *soap.Request
	var vars = soap.SubVars{PullPt: c.PullPt, TopicName: topic}

	// Pull point should already exist at this stage
	if c.PullPt != "" {
		soapReq, err = soap.NewRequest(soap.SUBSCRIBEPULL, vars)
		if err != nil {
			return ErrCantSubscribeTopicPull
		}

		var answer []byte

		answer, err = soapReq.Send(c.target)
		if err != nil {
			var res *STPPBody

			err = xml.Unmarshal(answer, res)
			if err != nil {
				c.List[topic].UnsubAddr = res.SubscribeTopicResponse.Reference.Address
				c.List[topic].Started = true
			} else {
				return
			}
		}
	}
	return
}
