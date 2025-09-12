package valueobject

import "testing"

func TestNewPort_Valid(t *testing.T) {
	portNum := 22
	port, err := NewPort(portNum)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if int(port) != portNum {
		t.Errorf("expected %d, got %d", portNum, port)
	}
}

func TestNewPort_Invalid(t *testing.T) {
	_, err := NewPort(0)
	if err == nil {
		t.Fatal("expected error for invalid port, got nil")
	}
	_, err = NewPort(70000)
	if err == nil {
		t.Fatal("expected error for invalid port, got nil")
	}
}
