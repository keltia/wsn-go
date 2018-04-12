// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"wsn-go/config"
	"wsn-ng/wsn"
)

func main() {
	config, err := config.LoadConfig("surveillance")
/*	pull := wsn.NewPullClient()
	err = pull.Subscribe("foo")
	defer pull.Stop()

	fmt.Printf("pull is of type: %v\n", reflect.TypeOf(pull))
*/
	push := wsn.NewPushClient(config)
	defer push.Stop()

	err = push.Subscribe("bar")

	fmt.Printf("push is of type: %v\n", reflect.TypeOf(push))


	push.Start()
//	pull.Start()
	data, err := ioutil.ReadAll(push)
	if err == nil {
		fmt.Println(string(data))
	}

	err = push.Unsubscribe("toto")
	if err != nil {
		fmt.Printf("Error unsubcribing %s: %v\n", "toto", err)
	}
}
