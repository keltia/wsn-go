// fa-export.go
//
//

package main

import (
	"../config"
	"../surv"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"text/tabwriter"
)

var (
	RcFile = filepath.Join(os.Getenv("HOME"), ".surveillance", "config.yml")

	Feeds = map[string]string{
		"AsterixJSON": "feed_json",
		"AsterixXML": "feed_xml",
		"AsterixJSONgzipped": "feed_jsongz",
	}
)

func main() {
	var feeds []string

	c, err := config.LoadConfig(RcFile)
	if err != nil {
		panic("Error loading "+RcFile)
	}
	fmt.Println(c.Dests)
	fmt.Println(c.Default, c.Dests[c.Default])

	flag.Parse()
	client, err := surv.NewClient(c)

	// Look for feed names on CLI
	for _, tn := range flag.Args {
		if Feeds[tn] != "" {
			feeds = append(feeds, tn)
		}
	}
	for _, tn := range feeds {
		// server := NewServer(tn
		unsubFn, err := client.Subscribe(tn, Feeds[tn])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error subscribing to %n: %v", tn, err)
		}
		topic := surv.Topic{Started: true, Address: unsubFn}
		client.Topics[tn] = topic
	}
	for name, topic := range(client.Topics) {
		fmt.Printf("Topic: %s Bytes: %ld Pkts: %d", name, topic.Bytes, topic.Pkts)
	}
}
