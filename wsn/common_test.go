package wsn

import (
	"testing"
	"github.com/keltia/wsn-go/config"
)

var common_config = config.Config{Proto: "http", Site: "example.com", Port: 666, Endpoint: "foo"}

func TestGenerateEndpoint(t *testing.T) {
	url := generateEndpoint(&common_config)
	if url != "http://example.com:666/foo" {
		t.Errorf("Error: bad format %s for %v\n", url, common_config)
	}
}

