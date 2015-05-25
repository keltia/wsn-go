package surv

import (
	"../config"
	"text/template"
	"os"
	"fmt"
	"runtime"
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
	err := t.Execute(os.Stdout, subvars)
	if err != nil {

	}
	topic := cl.Topics[name]
	topic.Started = true
	return "", nil
}

func (cl *Client) Unsubscribe(name string) (Topic, error) {
	topic := cl.Topics[name]
	if topic.Started {
		topic.Started = false
		return Topic{}, nil
	}
	return Topic{}, runtime.Error("Topic "+name+" not started.")
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