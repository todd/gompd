package mpd_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"
	"strconv"
	"testing"
)

var host = "localhost"
var port = 0
var timeoutms = 0

func TestGompd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gompd Suite")
}

var _ = BeforeSuite(func() {
	if env_host := os.Getenv("GOMPD_HOST"); env_host != "" {
		host = env_host
	}

	if env_port := os.Getenv("GOMPD_PORT"); env_port != "" {
		port, _ = strconv.Atoi(env_port)
	}

	if env_timeout_ms := os.Getenv("GOMPD_TIMEOUTMS"); env_timeout_ms != "" {
		timeoutms, _ = strconv.Atoi(env_timeout_ms)
	}
})
