package surv

import (
	"../config"
	"text/template"
	"os"
	"fmt"
)

var (
	Feeds = map[string]string{
		"AsterixJSON": "feed_json",
		"AsterixXML": "feed_xml",
		"AsterixJSONgzipped": "feed_jsongz",
	}
)

const (
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

func (cl *Client) Subscribe(name, callback string) (string, error) {
	c := cl.config
	subvars := SubVars{c.Base+":"+c.Port+"/"+callback, name}

	t := template.Must(template.New("subscribe").Parse(subText))
	if err := t.Execute(os.Stdout, subvars); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating template\n")
		return "", err
	}
	topic := cl.Topics[name]
	topic.Started = true
	return "", nil
}

func (cl *Client) Unsubscribe(name string) (Topic, error) {
	topic := cl.Topics[name]
	topic.Started = false
	return topic, nil
}

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