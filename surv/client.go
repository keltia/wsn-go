package surv

import (
	"../config"
	"text/template"
	"os"
	"fmt"
	"bytes"
	"net/http"
//	"io/ioutil"
)

var (
	Feeds = map[string]string{
		"AsterixJSON": "feed_json",
		"AsterixXML": "feed_xml",
		"AsterixJSONgzipped": "feed_jsongz",
	}
	survClient	http.Client = http.Client{}

	subText = `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
	               xmlns:b="http://docs.oasis-open.org/wsn/b-2"
	               xmlns:add="http://www.w3.org/2005/08/addressing">
	   <soap:Header/>
	   <soap:Body>
	      <b:Subscribe>
	         <b:ConsumerReference>
	            <add:Address>{{.my_topic}}</add:Address>
	         </b:ConsumerReference>
	         <b:Filter>
	           <b:TopicExpression>
	             {{.topic}}
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
	cl.config	= c
	cl.Target	= c.Proto+"://"+c.Site+":"+c.Port+"/"+c.Endpoint
	for i, _ := range(Feeds) {
		fmt.Println("Configuring "+i)
		cl.Topics[i] = Topic{ 0, 0, "", false}
	}
	cl.Feed_one = defaultFeed
	return cl, nil
}

func init() {
	survClient = http.Client{}
}

// Subscribe to a given topic
func (cl *Client) Subscribe(name, callback string) (string, error) {
	var result	bytes.Buffer

	c := cl.config
	subvars := SubVars{c.Base+":"+c.Port+"/"+callback, name}

	t := template.Must(template.New("subscribe").Parse(string(subText)))
	if err := t.Execute(&result, subvars); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating template\n")
		return "", err
	}

	targetURL := c.Base+":"+c.Port+"/"+callback
	payload := result.String()

	buf := bytes.NewBufferString(payload)
	req, err := http.NewRequest("POST", targetURL, buf)
	req.Header.Add("SOAPAction", "Subscribe")
	req.Header.Add("Content-Type", "text/xml; charset=UTF-8")

	resp, err := survClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during POST: %v", err)
		return "", nil
	}

	//body, err := ioutil.ReadAll(resp.Body)


	topic := cl.Topics[name]
	topic.Started = true
	return "", nil
}

// Unsubscribe
func (cl *Client) Unsubscribe(name string) (Topic, error) {
	topic := cl.Topics[name]
	buf := bytes.NewBufferString(unsubText)
	req, err := http.NewRequest("POST", topic.Address, buf)
	if err != nil {
		return Topic{}, err
	}
	req.Header.Add("SOAPAction", "Unsubscribe")
	req.Header.Add("Content-Type", "text/xml; charset=UTF-8")

	resp, err := survClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return Topic{}, err
	} else {
		topic.Started = false
		return topic, nil
	}
}

