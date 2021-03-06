package wsn

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"testing"
)

func TestCheckName(t *testing.T) {
	os.Setenv("HOME", "/home/foo")

	// Check tag usage
	file := "mytag"
	res := checkName(file)
	real := path.Join(os.Getenv("HOME"), fmt.Sprintf(".%s", file), "config.toml")
	if res != real {
		t.Errorf("Error: badly formed fullname %s—%s", res, real)
	}

	// Check fullname usage
	file = "/nonexistent/foobar.toml"
	res = checkName(file)
	if res != file {
		t.Errorf("Error: badly formed fullname %s", res)
	}

	// Check bad usage
	file = "/toto.yaml"
	res = checkName(file)
	if res != "" {
		t.Errorf("Error: should end with .toml: %s", res)
	}
}

func TestStringer(t *testing.T) {
	dest := Dest{Broker: "broker", Name: "myname", Type: "mytype"}

	res := dest.String()
	if res != "broker: myname" {
		t.Errorf("Error: malformed string: %s", res)
	}
}

func TestLoadConfig(t *testing.T) {
	file := "config.toml"
	conf, err := LoadConfig(file)
	if err != nil {
		t.Errorf("Malformed file %s: %v", file, err)
	}

	base := "http://147.196.152.4"
	if conf.Base != base {
		t.Errorf("Malformed base %s: %s", conf.Base, base)

	}

	site := "http://192.70.89.113"
	if conf.Target != site {
		t.Errorf("Malformed site %s: %s", conf.Target, site)
	}

	port := 9000
	if conf.Port != port {
		t.Errorf("Malformed port %d: %d", conf.Port, port)
	}

	endpoint := "wsn/NotificationBroker"
	if conf.Broker != endpoint {
		t.Errorf("Malformed base %s: %s", conf.Broker, endpoint)
	}

	def := "mine"
	if conf.Default != def {
		t.Errorf("Malformed default %s: %s", conf.Default, def)
	}

	vers := 1
	if conf.Version != vers {
		t.Errorf("Configuration file too old: %d", conf.Version)
	}
}

func TestLoadConfigDest(t *testing.T) {
	file := "config.toml"
	conf, err := LoadConfig(file)
	if err != nil {
		t.Errorf("Malformed file %s: %v", file, err)
	}

	// Check Dest
	if len(conf.Dests) != 2 {
		t.Errorf("Error loading Dests map[]: %v", conf.Dests)
	}

	if reflect.TypeOf(conf.Dests) != reflect.TypeOf(map[string]Dest{}) {
		t.Errorf("Error loading Dests map[]: wrong type %v—%s", conf.Dests, reflect.TypeOf(conf.Dests))
	}

	dst := conf.Dests[conf.Default]
	if reflect.TypeOf(dst) != reflect.TypeOf(Dest{}) {
		t.Errorf("Error loading Dests map[]: wrong type %v—%s", dst, reflect.TypeOf(Dest{}))
	}

	real := Dest{Broker: "localhost", Name: "surv", Type: "queue"}
	if dst != real {
		t.Errorf("Error loading Dests map[]: wrong name %v—%v", dst, real)
	}
}
