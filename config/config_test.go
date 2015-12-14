package config

import (
	"testing"
	"os"
	"path"
	"fmt"
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