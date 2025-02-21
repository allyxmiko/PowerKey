package action

import "testing"

func TestShutdown(t *testing.T) {
	err := Shutdown("winrm", "806906623", "192.168.1.240", 60)
	if err != nil {
		t.Error(err)
	}
}
