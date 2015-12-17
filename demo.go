package main

import (
	"wsn-ng/wsn"
	"fmt"
	"io/ioutil"
	"reflect"
)

func main() {
	pull := wsn.NewPullClient()
	err := pull.Subscribe("foo")
	defer pull.Stop()

	fmt.Printf("pull is of type: %v\n", reflect.TypeOf(pull))

	push := wsn.NewPushClient()
	defer push.Stop()
	err = push.Subscribe("bar")
	err = push.Subscribe("baz")

	fmt.Printf("push is of type: %v\n", reflect.TypeOf(push))


	push.Start()
	pull.Start()
	data, err := ioutil.ReadAll(pull)
	if err == nil {
		fmt.Println(string(data))
	}

	data, err = ioutil.ReadAll(push)
	if err == nil {
		fmt.Println(string(data))
	}

	err = push.Unsubscribe("toto")
	if err != nil {
		fmt.Printf("Error unsubcribing %s: %v\n", "toto", err)
	}
}
