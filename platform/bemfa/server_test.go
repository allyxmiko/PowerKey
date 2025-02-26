package bemfa

import (
	"testing"
)

func TestServer(t *testing.T) {
	server, _ := NewServer("2e8c91d03a0e4a0a9f43a57329fd85bb")
	server.Subscription("pc001")
}
