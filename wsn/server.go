// server.go
//
// Server-side of the WS-N package

package wsn

import (
	"net/http"
	"encoding/xml"
	"fmt"
	"log"
	"io/ioutil"
	"strings"
)

func makeHandler(fn func(http.ResponseWriter, *http.Request, string, *Client), cl *Client) http.HandlerFunc {
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

func handleNotify(w http.ResponseWriter, req *http.Request, url string, cl *Client) {
	// body is an XML SOAP
	//
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()

		if cl.Verbose {
			log.Println("In handleNotify")
			log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)
			log.Println(req)

			log.Printf("|%s|\n", string(body))
		}

		last := getFeedName(url)
		if cl.Verbose {
			log.Println("Request is", last)
		}

		notify := &SurvData{}
		err = xml.Unmarshal(body, notify)
		if err != nil {
			real := fmt.Sprintf("Error: %v", err)
			log.Println(real)
			//http.Error(w, real, 500)
			return
		}

		topic := cl.Topics[last]
		topic.Bytes += int64(len(notify.Body.Notify.NotifyMsg.Message))
		topic.Pkts++

		(cl.Feed_one)(notify.Body.Notify.NotifyMsg.Message)
	} else {
		http.NotFound(w, req)
	}
}

func (cl *Client) ServerStart(feeds map[string]string) {
	server := http.NewServeMux()
	for name, feed := range feeds {
		log.Printf("Setting handler for %s as /%s\n", name, feed)
		server.HandleFunc("/" + feed, makeHandler(handleNotify, cl))
	}
	log.Println("Serving", feeds)
	log.Fatal(http.ListenAndServe(":"+fmt.Sprintf("%d", cl.Config.Port), server))
}
