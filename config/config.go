// config.go
//
// My homemade configuration class

package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/naoina/toml"
)

type Dest struct {
	Broker	string
	Name	string
	Type	string
}

type Config struct {
	Proto		string			// http
	Site 		string			// 192.70.89.113
	Port		int				// 9000
	Endpoint	string			// wsn/NotificationBroker
	Wsdl		string
	Base		string			// http://147.196.152.4
	Dests		map[string]Dest
	Default		string			// mine
}

// Check the parameter for either tag or filename
func checkName(file string) string {
	// Check for tag
	if !strings.HasSuffix(file, ".toml") {
		// file must be a tag so add a "."
		return filepath.Join(os.Getenv("HOME"),
			fmt.Sprintf(".%s", file),
			"config.toml")
	} else {
		return file
	}
}

// Basic Stringer for Config
func (dest Dest) String() string {
	return fmt.Sprintf("%v: %v", dest.Broker, dest.Name)
}

// Load a file as a TOML document and return the structure
func LoadConfig(file string) (*Config, error) {
	// Check for tag
	sFile := checkName(file)

	c := new(Config)
	buf, err := ioutil.ReadFile(sFile)
	if err != nil {
		return c, errors.New(fmt.Sprintf("Can not read %s", sFile))
	}

	err = toml.Unmarshal(buf, &c)
	if err != nil {
		return c, errors.New(fmt.Sprintf("Error parsing toml %s: %v",
			sFile, err))
	}

	// Finally set default destination
	c.Default = "mine"

	return c, nil
}
