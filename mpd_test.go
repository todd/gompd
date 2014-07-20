package mpd

import (
	"testing"
	"os"
	"strconv"
)

var host = "localhost"
var port = 0
var timeoutms = 0

func init() {
	if env_host := os.Getenv("GOMPD_HOST"); env_host != "" {
		host = env_host
	}

	if env_port := os.Getenv("GOMPD_PORT"); env_port != "" {
		port, _ = strconv.Atoi(env_port)
	}

	if env_timeout_ms := os.Getenv("GOMPD_TIMEOUTMS"); env_timeout_ms != "" {
		timeoutms, _ = strconv.Atoi(env_timeout_ms)
	}
}

func TestClientInit(t *testing.T) {
	if client, err := Init(host, port, timeoutms); err != nil {
		t.Errorf("Init failed: %v", err)
	} else {
		client.Close()
	}
}
