// config.go
//
// My homemade configuration class

package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Dest struct {
	Broker	string
	Name	string
	Type	string
}

type Config struct {
	Proto		string			// http
	Site 		string			// 192.70.89.113
	Port		string			// 9000
	Endpoint	string			// wsn/NotificationBroker
	Wsdl		string
	Base		string			// http://147.196.152.4
	Dests		map[string]Dest
	Default		string			// mine
}

func (dest Dest) String() string {
	return fmt.Sprintf("%v: %v", dest.Broker, dest.Name)
}

func LoadConfig(file string) (Config, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return Config{}, err
	}

	c := new(Config)
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		fmt.Println("Error parsing yaml")
		return Config{}, err
	} else {
		c.Default = "mine"
		return *c, err
	}
}
