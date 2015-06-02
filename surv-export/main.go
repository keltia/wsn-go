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
	"log"
	"os/signal"
	"time"
)

var (
	RcFile = filepath.Join(os.Getenv("HOME"), ".surveillance", "config.yml")

	// All possible feeds
	Feeds = map[string]string{
		"AsterixJSON": "feed_json",
		"AsterixXML": "feed_xml",
		"AsterixJSONgzipped": "feed_jsongz",
	}

	RunningFeeds = map[string]string{}

	SurvClient	*surv.Client
)

// Subscribe to wanted topics
func doSubscribe(feeds []string) {
	time.Sleep(5)
	// Go go go
	for _, tn := range feeds {
		unsubFn, err := SurvClient.Subscribe(tn, Feeds[tn])
		if err != nil {
			log.Printf("Error subscribing to %n: %v", tn, err)
		}
		if fVerbose {
			fmt.Println("Subscribing to ", tn, " as ", Feeds[tn])
			fmt.Println("  unsub is ", unsubFn)
		}
		topic := surv.Topic{Started: true, Address: unsubFn}
		SurvClient.Topics[tn] = topic
	}
}

// Handle shutdown operations
func doShutdown() {
	// do last actions and wait for all write operations to end
	for name, topic := range (SurvClient.Topics) {
		if topic.Started {
			err := SurvClient.Unsubscribe(name)
			if err != nil {
				log.Printf("Error unsubscribing to %n: %v", name, err)
			}
			if fVerbose {
				log.Println("Unsubscribing from", name)
				log.Printf("Topic: %s Bytes: %d Pkts: %d", name, topic.Bytes, topic.Pkts)
			}
		}
	}
}

// return list of keys of map m
func keys(m map[string]string) []string {
	var keys []string
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

// Main program
func main() {
	// Handle SIGINT
	go func() {
	    sigint := make(chan os.Signal, 3)
	    signal.Notify(sigint, os.Interrupt)
	    <-sigint
	    log.Println("Program killed !")

		doShutdown()

	    os.Exit(0)
	}()

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
			if fVerbose {
				log.Println("Configuring "+tn)
			}
			RunningFeeds[tn] = Feeds[tn]
			SurvClient.AddFeed(tn)
		}
	}

	// Start server for callback
	log.Println("Starting server for ", keys(RunningFeeds), "...")
	go doSubscribe(RunningFeeds)
	surv.ServerStart(SurvClient, RunningFeeds)
}
