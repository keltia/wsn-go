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

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
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
		fn(w, r, title)
	}
}

// Get the last component of URI
func getFeedName(url string) string {
	// FIXME
	parts := strings.Split(url, "/")
	return parts[len(parts) - 1]

}

func handleNotify(w http.ResponseWriter, req *http.Request, url string) {
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)
	log.Println(req)
	//
	// body is an XML SOAP
	//
	if req.Method == "POST" {
		log.Println("In handleNotify")
		body, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()

		last := getFeedName(url)
		log.Println("Request is", last)

		notify := &SurvData{}
		err = xml.Unmarshal(body, notify)
		if err != nil {
			real := fmt.Sprintf("Error: %v", err)
			log.Println(real)
			//http.Error(w, real, 500)
			return
		}

		//topicList := survClient.Topics
		//log.Println(topicList[last])

		//topicList[last].Bytes += int64(len(notify.Body.Notify.Message.Payload.Data))
		//topicList[last].Pkts++

		log.Printf("%+v", string(notify.Body.Notify.NotifyMsg.Message.Payload.Text))
	} else {
		http.NotFound(w, req)
	}
}

func ServerStart(cl *Client, feeds map[string]string) {
	server := http.NewServeMux()
	for name, feed := range feeds {
		log.Printf("Setting handler for %s as /%s\n", name, feed)
		server.HandleFunc("/" + feed, makeHandler(handleNotify))
	}
	log.Println("Serving", feeds)
	log.Fatal(http.ListenAndServe(":"+cl.Config.Port, server))
}
