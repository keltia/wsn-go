// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package main

import (
	"fmt"
	"github.com/keltia/wsn-go/config"
	"github.com/keltia/wsn-go/wsn"
	"os"
	"reflect"
	"log"
)

func main() {
	var buf = make([]byte, 262144)

	config, err := config.LoadConfig("surveillance")
	if err != nil {
		fmt.Printf("%v\n%v\n", config, err)
		os.Exit(1)
	}
	/*	pull := wsn.NewPullClient()
		err = pull.Subscribe("foo")
		defer pull.Stop()

		fmt.Printf("pull is of type: %v\n", reflect.TypeOf(pull))
	*/
	push := wsn.NewPushClient(config)
	defer push.Stop()

	err = push.Subscribe("bar")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("push is of type: %v\n", reflect.TypeOf(push))

	push.SetTimeout(10)
	o := make(chan []byte, 262144)
	push.Start(o)
	//	pull.Start()
	for {
		buf = <-o
		fmt.Println(string(buf))

		if buf == nil {
			break
		}
	}

	err = push.Unsubscribe("toto")
	if err != nil {
		fmt.Printf("Error unsubcribing %s: %v\n", "toto", err)
	}
}
