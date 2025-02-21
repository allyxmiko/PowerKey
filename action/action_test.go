package action

import "testing"

func TestShutdown(t *testing.T) {
	t.Skip()
	err := Shutdown("winrm", "806906623", "192.168.1.240", 60)
	if err != nil {
		t.Error(err)
	}
}

func TestWakeOnLAN(t *testing.T) {
	t.Skip()
	err := WakeOnLAN("4C-10-D5-81-CF-EF")
	if err != nil {
		t.Error(err)
	}
}

func TestPing(t *testing.T) {
	t.Skip()
	err := Ping("192.168.1.240")
	if err != nil {
		t.Error(err)
	}
}
