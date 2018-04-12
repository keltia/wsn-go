// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn_test

import (
	"fmt"
	"reflect"
	"io/ioutil"
	"wsn-ng/wsn"
)

func ExampleNewPullClient() {
	pull := wsn.NewPullClient()
	err := pull.Subscribe("foo")
	defer pull.Stop()

	fmt.Printf("pull is of type: %v\n", reflect.TypeOf(pull))

	pull.Start()

	data, err := ioutil.ReadAll(pull)
	if err == nil {
		fmt.Println(string(data))
	}

	data, err = ioutil.ReadAll(pull)

	pull.Stop()
}
