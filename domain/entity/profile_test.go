package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProfile_Valid(t *testing.T) {
	id := 1
	name := "example.com"
	username := "user"
	ip := "192.168.1.1"
	port := 22
	group := "test"
	details := []string{"detail1", "detail2"}

	profile, err := NewProfile(id, name, username, ip, port, group, details)

	assert.NoError(t, err)
	assert.NotNil(t, profile)
	assert.Equal(t, id, profile.ID)
	assert.Equal(t, group, profile.Group)
	assert.Equal(t, ip, string(profile.IP))
	assert.Equal(t, name, string(profile.Name))
	assert.Equal(t, port, int(profile.Port))
	assert.Equal(t, username, string(profile.Username))
	assert.Equal(t, details, profile.Details)
}

func TestNewProfile_InvalidIP(t *testing.T) {
	_, err := NewProfile(1, "example.com", "user", "invalid-ip", 22, "test", []string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid IP address")
}

func TestNewProfile_InvalidPort(t *testing.T) {
	_, err := NewProfile(1, "example.com", "user", "192.168.1.1", 0, "test", []string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid port")
}

func TestProfile_GetFullAddress(t *testing.T) {
	profile, _ := NewProfile(1, "example.com", "user", "192.168.1.1", 22, "test", []string{})
	expected := "192.168.1.1:22"
	assert.Equal(t, expected, profile.GetFullAddress())
}

func TestProfile_HasDetails(t *testing.T) {
	tests := []struct {
		name    string
		details []string
		want    bool
	}{
		{
			name:    "has details",
			details: []string{"detail"},
			want:    true,
		},
		{
			name:    "no details",
			details: []string{},
			want:    false,
		},
		{
			name:    "nil details",
			details: nil,
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			profile, _ := NewProfile(1, "example.com", "user", "192.168.1.1", 22, "test", tt.details)
			assert.Equal(t, tt.want, profile.HasDetails())
		})
	}
}

func TestProfile_IsSame(t *testing.T) {
	profile1, _ := NewProfile(1, "example.com", "user", "192.168.1.1", 22, "test", []string{})
	profile2, _ := NewProfile(2, "different.com", "user", "192.168.1.1", 22, "other", []string{})
	profile3, _ := NewProfile(3, "example.com", "different", "192.168.1.1", 22, "test", []string{})
	profile4, _ := NewProfile(4, "example.com", "user", "192.168.1.2", 22, "test", []string{})

	assert.True(t, profile1.IsSame(profile2))  // same user and IP
	assert.False(t, profile1.IsSame(profile3)) // different user
	assert.False(t, profile1.IsSame(profile4)) // different IP
}
