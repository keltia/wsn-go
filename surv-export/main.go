// fa-export.go
//
//

package main

import (
	"../config"
	"../surv"
	"fmt"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
)

var (
	RcFile = filepath.Join(os.Getenv("HOME"), ".surveillance", "config.yml")
)

func main() {
	app := cli.NewApp()
	app.Name = "surv-export"
	app.Author = "Ollivier Robert"
	app.Version = "0.0.1"
	app.Usage = "surv-export"

	c, err := config.LoadConfig(RcFile)
	if err != nil {
		fmt.Println("Error loading")
	}
	fmt.Println(c.Dests)
	fmt.Println(c.Default, c.Dests[c.Default])

	client, err := surv.NewClient(c)
	for name, topic := range(client.Topics) {
		fmt.Printf("Topic: %s Bytes: %ld Pkts: %d", name, topic.Bytes, topic.Pkts)
	}
}
