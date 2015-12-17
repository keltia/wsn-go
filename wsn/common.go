// common.go

/*
  Package which implements some common routines to the Push & Pull clients
 */
package wsn

import (
	"fmt"
	"net/http"
	"github.com/keltia/wsn-go/config"
)

var (
	httpClient http.Client = http.Client{}

	subText = `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
	               xmlns:b="http://docs.oasis-open.org/wsn/b-2"
	               xmlns:add="http://www.w3.org/2005/08/addressing">
	   <soap:Header/>
	   <soap:Body>
	      <b:Subscribe>
	         <b:ConsumerReference>
	            <add:Address>{{.TopicURL}}</add:Address>
	         </b:ConsumerReference>
	         <b:Filter>
	           <b:TopicExpression>
	             {{.TopicName}}
	           </b:TopicExpression>
	         </b:Filter>
	      </b:Subscribe>
	   </soap:Body>
	</soap:Envelope>
	`

	unsubText = `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
	               xmlns:b="http://docs.oasis-open.org/wsn/b-2">
	   <soap:Header/>
	   <soap:Body>
	      <b:Unsubscribe/>
	   </soap:Body>
	</soap:Envelope>
	`
)


type Client interface {
	Subscribe(string, string) (string, error)
	Unsubscribe(string) error
	AddFeed(string)
	AddHandler(func([]byte))
	SetTime(int64)
}

// Generate an URL
func generateEndpoint(c *config.Config) string {
	return fmt.Sprintf("%s://%s:%d/%s", c.Proto, c.Site, c.Port, c.Endpoint)
}


// Defaults to console output
func defaultFeed(buf []byte) { fmt.Println(string(buf))}

