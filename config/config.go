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
	Proto		string
	Site 		string
	Port		string
	Endpoint	string
	Wsdl		string
	Base		string
	Dests		map[string]Dest
	Default		string
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
	}

	c.Default = "mine"
	return *c, err
}
