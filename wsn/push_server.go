// Copyright 2015 Ollivier Robert for EUROCONTROL  All rights reserved

package wsn

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func makeHandler(fn func(http.ResponseWriter, *http.Request, string, *PushClient), client *PushClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//defer func() {
		//	if e, ok := recover().(error); ok {
		//		log.Println(e)
		//		http.Error(w, e.Error(), http.StatusInternalServerError)
		//		return
		//	}
		//}()
		title := r.URL.Path
		log.Println(title)
		fn(w, r, title, client)
	}
}

// Get the last component of URI
func getFeedName(url string) string {
	// FIXME
	parts := strings.Split(url, "/")
	return parts[len(parts) - 1]

}

func handleNotify(w http.ResponseWriter, req *http.Request, url string, client *PushClient) {
	// body is an XML SOAP
	//
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body);
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading payload %s\n", req.Body), 500)
		}
		defer req.Body.Close()

		if client.verbose {
			log.Println("In handleNotify")
			log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)
			log.Println(req)
		}

		last := getFeedName(url)
		if client.verbose {
			log.Println("Request is", last)
		}

		// Decode body
		notify := &WsnData{}
		err = xml.Unmarshal(body, notify)
		if err != nil {
			log.Println(fmt.Sprintf("Error parsing: |%s|: %v", string(body), err))
			//http.Error(w, real, 500)
		}

		// payload is way inside
		payload := notify.Body.Notify.NotificationMessage.Message.Data

		// XXX very long output, convert to string for clarity
		if client.verbose {
			log.Printf("|%v|", string(payload))
		}

		// Stats
		topic := client.List[last]
		topic.Bytes += int64(len(payload))
		topic.Pkts++

		// Send everything through the channel
		client.output<- payload
	} else {
		http.NotFound(w, req)
	}
}

func (client *PushClient) StartServer() {
	server := http.NewServeMux()
	for feed, _ := range client.List {
		log.Printf("Setting handler %s\n", feed)
		server.HandleFunc(feed, makeHandler(handleNotify, client))
	}
	log.Println("Starting")
	log.Fatal(http.ListenAndServe(":"+fmt.Sprintf("%d", client.port), server))
}

