package surv

import (
	"../config"
	"text/template"
	"log"
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"strings"
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

func defaultFeed(buf []byte) { fmt.Println(buf)}

func NewClient (c config.Config) (*Client, error) {
	cl := new(Client)
	cl.Topics	= make(map[string]Topic, 10)
	cl.Config	= c
	cl.Target	= c.Proto+"://"+c.Site+":"+c.Port+"/"+c.Endpoint
	cl.Feed_one = defaultFeed
	return cl, nil
}

// Create .Topics structure w/o subscribing
func (cl *Client) AddFeed(name string) {
	log.Println("Adding new feed", name)
	cl.Topics[name] = Topic{Started: false}
}

// Generate an URL
func (cl *Client) generateURL() string {
	c := cl.Config
	return fmt.Sprintf("%s://%s:%s/%s", c.Proto, c.Site, c.Port, c.Endpoint)
}

// Subscribe to a given topic
func (cl *Client) Subscribe(name, callback string) (string, error) {
	var result	bytes.Buffer

	c := cl.Config
	targetURL := cl.generateURL()
	myEndpoint := cl.Config.Base + ":" + cl.Config.Port + "/" + callback

	log.Println("Targetting ", targetURL)
	log.Printf("Subscribing %s on my side", myEndpoint)

	subvars := SubVars{TopicURL: myEndpoint, TopicName: name}

	t := template.Must(template.New("subscribe").Parse(string(subText)))
	if err := t.Execute(&result, subvars); err != nil {
		log.Printf("Error creating template\n")
		return "", err
	}

	payload := result.String()

	buf := bytes.NewBufferString(payload)
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
	req.Header.Add("SOAPAction", "Unsubscribe")
	req.Header.Add("Content-Type", "text/xml; charset=UTF-8")

	resp, err := httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return err
	} else {
		topic.Started = false
		return nil
	}
}

