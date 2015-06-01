// server.go
//
// Server-side of the WS-N package

package surv

import (
	"net/http"
	"encoding/xml"
	"fmt"
	"log"
	"io/ioutil"
	"strings"
)

var survClient *Client

func handleNotify(w http.ResponseWriter, req *http.Request) {
	//
	// body is an XML SOAP
	//
	log.Println("In handleNotify")
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	notify := SurvData{}
	err = xml.Unmarshal(body, &notify)
	if err != nil {
		real := fmt.Sprintf("Error: %v", err)
		http.Error(w, real, 500)
	}
	// FIXME
	pathInfo := req.Header.Get("Path-Info")
	log.Println(pathInfo)
	parts := strings.Split(pathInfo, "/")
	last := parts[len(parts) - 1]
	log.Println(last)

	topicList := survClient.Topics
	log.Println(topicList[last])

	//topicList[last].Bytes += int64(len(notify.Body.Notify.Message.Payload.Data))
	//topicList[last].Pkts++

	fmt.Println(notify.Body.Notify.Message.Payload.Data)
}

func ServerStart(cl *Client, feeds []string) {
	survClient = cl
	server := http.NewServeMux()
	for _, feed := range feeds {
		log.Println("Setting handler for "+feed)
		server.HandleFunc(feed, handleNotify)
	}
	log.Println("Serving...")
	log.Fatal(http.ListenAndServe(":"+cl.config.Port, nil))
}
