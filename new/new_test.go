package new

import "testing"

func TestNewService(t *testing.T) {
	server := NewService()
	server.Start()
	server.Stop()
}