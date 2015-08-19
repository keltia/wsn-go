// surv-export.go
//
// Export data from the WS-N endpoint giving out ADS-B data in various formats.
//
// @author Ollivier Robert <ollivier.robert@eurocontrol.int>

package main

import (
	"../config"
	"../surv"
	"flag"
	"fmt"
	"os"
	"log"
	"os/signal"
	"time"
	"strconv"
	"regexp"
)

var (
	// We use a tag to find the proper file now. $HOME/.<tag>/config.toml
	RcFile = "surveillance"

	// All possible feeds
	Feeds = map[string]string{
		"AsterixJSON": "feed_json",
		"AsterixXML": "feed_xml",
		"AsterixJSONgzipped": "feed_jsongz",
	}

	timeMods = map[string]int64{
		"mn": 60,
		"h":  3600,
		"d":  3600 * 24,
	}

	RunningFeeds = map[string]string{}

	SurvClient	*surv.Client

	fOutputFH	*os.File
)

// Subscribe to wanted topics
func doSubscribe(feeds map[string]string) {
	time.Sleep(1 * time.Second)
	// Go go go
	for name, target := range feeds {
		unsubFn, err := SurvClient.Subscribe(name, target)
		if err != nil {
			log.Printf("Error subscribing to %n: %v", name, err)
		}
		if fVerbose {
			log.Printf("Subscribing to /%s for %s\n", target, name)
			log.Printf("  unsub is %s\n", unsubFn)
		}
		topic := surv.Topic{Started: true, UnsubAddr: unsubFn}
		SurvClient.Topics[name] = topic
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

// fOutput callback
func fileOutput(buf []byte) {
	if nb, err := fOutputFH.Write(buf); err != nil {
		log.Fatalf("Error writing %d bytes: %v", nb, err)
	}
}

// Check for specific modifiers, returns seconds
//
//XXX could use time.ParseDuration except it does not support days.
func checkTimeout(value string) int64 {
	mod := int64(1)
	re := regexp.MustCompile(`(?P<time>\d+)(?P<mod>(s|mn|h|d)*)`)
	match := re.FindStringSubmatch(value)
	if match == nil {
		return 0
	} else {
		// Get the base time
		time, err := strconv.ParseInt(match[1], 10, 64)
		if err != nil {
			return 0
		}

		// Look for meaningful modifier
		if match[2] != "" {
			mod = timeMods[match[2]]
			if mod == 0 {
				mod = 1
			}
		}

		// At the worst, mod == 1.
		return time * mod
	}
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

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Fprint(os.Stderr, "You must specify at least one feed!\n")
		fmt.Fprintln(os.Stderr, "List of possible feeds:")
		for f, _ := range Feeds {
			fmt.Fprintf(os.Stderr, "  %s\n", f)
		}
		os.Exit(1)
	}

	// Load configuration
	c, err := config.LoadConfig(RcFile)
	if err != nil {
		panic("Error loading "+RcFile)
	}
	if fVerbose {
		fmt.Printf("Config is %s://%s:%s/%s\n", c.Proto, c.Site, c.Port, c.Endpoint)
		fmt.Println(c.Dests)
		fmt.Println(c.Default, c.Dests[c.Default])
	}

	// Actually instanciate the client part
	if SurvClient, err = surv.NewClient(c); err != nil {
		log.Fatalf("Error connecting to %s: %v", SurvClient.Target)
	}

	if fVerbose {
		SurvClient.Verbose = true
	}

	// Open output file
	if (fOutput != "") {
		if (fVerbose) {
			log.Printf("Output file is %s\n", fOutput)
		}

		if fOutputFH, err = os.Create(fOutput); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating %s\n", fOutput)
			panic(err)
		}

		SurvClient.AddHandler(fileOutput)
	}

	// Check if we did specify a timeout with -i
	if fsTimeout != "" {
		fTimeout = checkTimeout(fsTimeout)

		if fVerbose {
			log.Printf("Run for %ds\n", fTimeout)
		}
		SurvClient.SetTimer(fTimeout)
	}

	// Look for feed names on CLI
	for _, tn := range flag.Args() {
		if Feeds[tn] != "" {
			if fVerbose {
				log.Println("Configuring", Feeds[tn], "for", tn)
			}
			RunningFeeds[tn] = Feeds[tn]
			SurvClient.AddFeed(tn)
		}
	}

	// Start server for callback
	log.Println("Starting server for ", keys(RunningFeeds), "...")
	go doSubscribe(RunningFeeds)
	SurvClient.ServerStart(RunningFeeds)
}
