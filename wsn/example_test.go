// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn_test

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"wsn-go/wsn"
)

func ExampleNewPullClient() {
	config := LoadConfig()
	pull := wsn.NewPullClient(config)
	err := pull.Subscribe("foo")
	defer pull.Stop()

	fmt.Printf("pull is of type: %v\n", reflect.TypeOf(pull))

	output := make(chan []byte, 262144)
	pull.Start(output)

	data, err := ioutil.ReadAll(pull)
	if err == nil {
		fmt.Println(string(data))
	}

	pull.Stop()
}
