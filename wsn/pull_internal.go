// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"encoding/xml"
	"wsn-go/soap"
)

// createPullPoint create a Pull point on the broker that will be used to subscribe
// topics.
func (c *PullClient) createPullPoint() (pullPt string, err error) {

	// Pull point should not exist at this stage
	if c.PullPt == "" {
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
