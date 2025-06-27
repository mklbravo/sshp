package valueobject

import (
	"testing"
)

func TestNewIP_Valid(t *testing.T) {
	ipStr := "192.168.1.1"
	ip, err := NewIP(ipStr)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if string(ip) != ipStr {
		t.Errorf("expected %s, got %s", ipStr, ip)
	}
}

func TestNewIP_Invalid(t *testing.T) {
	_, err := NewIP("invalid-ip")
	if err == nil {
		t.Fatal("expected error for invalid IP, got nil")
	}
}
