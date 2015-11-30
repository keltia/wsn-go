package wsn

import (
	"text/template"
	"log"
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"strings"
	"time"
	"os"
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

// Private functions

// Defaults to console output
func defaultFeed(buf []byte) { fmt.Println(string(buf))}

// Generate an URL
func (cl *Client) generateURL() string {
	c := cl.Config
	return fmt.Sprintf("%s://%s:%d/%s", c.Proto, c.Site, c.Port, c.Endpoint)
}

// Public interface

// Create new client instance
func NewClient (c *config.Config) (*Client, error) {
	cl := new(Client)
	cl.Topics	= make(map[string]Topic, 10)
	cl.Config	= c
	cl.Target	= c.Proto+"://"+c.Site+":"+fmt.Sprintf("%d", c.Port)+"/"+c.Endpoint
	cl.Feed_one = defaultFeed
	return cl, nil
}

// Create .Topics structure w/o subscribing
func (cl *Client) AddFeed(name string) {
	if cl.Verbose {
		log.Println("Adding new feed", name)
	}
	cl.Topics[name] = Topic{Started: false}
}

// Change default callback
func (cl *Client) AddHandler(fn func([]byte)) {
	cl.Feed_one = fn
}

// Allow run of specified duration
func (cl *Client) SetTimer(timer int64) {
	// Sleep for fTimeout seconds then sends Interrupt
	go func() {
		time.Sleep(time.Duration(timer) * time.Second)
		if cl.Verbose {
			log.Println("Timer off, time to kill")
		}
		myself, _ := os.FindProcess(os.Getpid())
		myself.Signal(os.Interrupt)
	}()
}

// Subscribe to a given topic
func (cl *Client) Subscribe(name, callback string) (string, error) {
	var result	bytes.Buffer

	c := cl.Config
	targetURL := cl.generateURL()
	myEndpoint := cl.Config.Base + ":" + fmt.Sprintf("%d", cl.Config.Port) + "/" + callback

	if cl.Verbose {
		log.Println("Targetting ", targetURL)
		log.Printf("Subscribing %s on my side", myEndpoint)
	}

	subvars := SubVars{TopicURL: myEndpoint, TopicName: name}

	t := template.Must(template.New("subscribe").Parse(string(subText)))
	if err := t.Execute(&result, subvars); err != nil {
		log.Printf("Error creating template\n")
		return "", err
	}

	buf := bytes.NewBufferString(result.String())
	req, err := http.NewRequest("POST", targetURL, buf)
	if err != nil {
		log.Fatal("Error creating request for ", buf, ": ", err)
	}
	req.Header.Set("SOAPAction", "Subscribe")
	req.Header.Set("Content-Type", "text/xml; charset=UTF-8")

	resp, err := httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.Printf("Error during POST: %v", err)
		return "", nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	// Parse XML
	res := &SubscribeAnswer{}
	if err = xml.Unmarshal(body, res); err != nil {
		return "", err
	}

	address := res.Body.Resp.Reference.Address
	address = strings.Replace(address, "0.0.0.0", c.Site, -1)

	topic := cl.Topics[callback]
	topic.Started = true
	topic.UnsubAddr = address

	return address, nil
}

// Unsubscribe
func (cl *Client) Unsubscribe(name string) (error) {
	topic := cl.Topics[name]
	buf := bytes.NewBufferString(unsubText)
	req, err := http.NewRequest("POST", topic.UnsubAddr, buf)
	if err != nil {
		return err
	}
	req.Header.Set("SOAPAction", "Unsubscribe")
	req.Header.Set("Content-Type", "text/xml; charset=UTF-8")

	resp, err := httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return err
	} else {
		topic.Started = false
		return nil
	}
}

