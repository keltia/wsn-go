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
)

var (
	RcFile = filepath.Join(os.Getenv("HOME"), ".surveillance", "config.yml")

	Feeds = map[string]string{
		"AsterixJSON": "feed_json",
		"AsterixXML": "feed_xml",
		"AsterixJSONgzipped": "feed_jsongz",
	}

	SurvClient	surv.Client
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
	SurvClient, err := surv.NewClient(c)

	// Look for feed names on CLI
	for _, tn := range flag.Args() {
		if Feeds[tn] != "" {
			feeds = append(feeds, tn)
			SurvClient.NewFeed(Feeds[tn])
			if fVerbose {
				fmt.Println("Configuring "+tn)
			}
		}
	}

	// Start server for callback
	// server := NewServer(tn

	// Go go go
	for _, tn := range feeds {
		unsubFn, err := SurvClient.Subscribe(tn, Feeds[tn])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error subscribing to %n: %v", tn, err)
		}
		topic := surv.Topic{Started: true, Address: unsubFn}
		SurvClient.Topics[tn] = topic
	}
	for name, topic := range(SurvClient.Topics) {
		fmt.Printf("Topic: %s Bytes: %ld Pkts: %d", name, topic.Bytes, topic.Pkts)
	}
}
