// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import "fmt"

// generateURL generates an URL on the target site
func (c *PullClient) generateURL(endPoint string) string {
	return fmt.Sprintf("%s:%d/%s", c.target, c.port, endPoint)
}

// createPullPoint create a Pull point on the broker that will be used to subscribe
// topics.
func createPullPoint() (pullPt string, err error) {
	if false {
		err = ErrCreatingPullPoint
	}
	return
}

// destroyPullPoint de-registers the pull point to avoid hogging resources on the broker
func destroyPullPoint(pullPt string) (err error) {
	if false {
		err = ErrDestroyingPullPoint
	}
	return
}
