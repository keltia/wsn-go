// push_server.go
//
// Server-side of the WS-N Push package

package wsn

import (
	"net/http"
	"encoding/xml"
	"fmt"
	"log"
	"io/ioutil"
	"strings"
)

func makeHandler(fn func(http.ResponseWriter, *http.Request, string, *PushClient), cl *PushClient) http.HandlerFunc {
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
		fn(w, r, title, cl)
	}
}

// Get the last component of URI
func getFeedName(url string) string {
	// FIXME
	parts := strings.Split(url, "/")
	return parts[len(parts) - 1]

}

func handleNotify(w http.ResponseWriter, req *http.Request, url string, cl *PushClient) {
	// body is an XML SOAP
	//
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body);
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading payload %s\n", req.Body), 500)
		}
		defer req.Body.Close()

		if cl.Verbose {
			log.Println("In handleNotify")
			log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)
			log.Println(req)
		}

		last := getFeedName(url)
		if cl.Verbose {
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
		if cl.Verbose {
			log.Printf("|%v|", string(payload))
		}

		// Stats
		topic := cl.Topics[last]
		topic.Bytes += int64(len(payload))
		topic.Pkts++

		(cl.Feed_one)(payload)
	} else {
		http.NotFound(w, req)
	}
}

func (cl *PushClient) ServerStart() {
	server := http.NewServeMux()
	for feed, _ := range cl.Topics {
		log.Printf("Setting handler %s\n", feed)
		server.HandleFunc(feed, makeHandler(handleNotify, cl))
	}
	log.Println("Starting")
	log.Fatal(http.ListenAndServe(":"+fmt.Sprintf("%d", cl.Config.Port), server))
}
