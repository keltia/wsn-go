// server.go
//
// Server-side of the WS-N package

package surv

import (
	"../config"
	"net/http"
	"encoding/xml"
	"fmt"
	"log"
	"io/ioutil"
)

func handleNotify(w http.ResponseWriter, req *http.Request) {
	//
	// body is an XML SOAP
	//
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	notify := SurvData{}
	err = xml.Unmarshal(body, &notify)
	if err != nil {
		real := fmt.Sprintf("Error: %v", err)
		http.Error(w, real, 500)
	}
	fmt.Println(notify.Body.Notify.Message.Payload.Data)
}

func ServerStart(cl *Client, feeds []string) {
	server := http.NewServeMux()
	for _, feed := range feeds {
		log.Println("Setting handler for "+feed)
		server.HandleFunc(feed, handleNotify)
	}
	log.Println("Serving...")
	log.Fatal(http.ListenAndServe(":"+cl.config.Port, nil))
}
