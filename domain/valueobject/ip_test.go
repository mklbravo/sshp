package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIP(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "valid IPv4",
			input:   "192.168.1.1",
			want:    "192.168.1.1",
			wantErr: false,
		},
		{
			name:    "valid IPv6",
			input:   "::1",
			want:    "::1",
			wantErr: false,
		},
		{
			name:    "invalid IP",
			input:   "invalid-ip",
			want:    "",
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "out of range IPv4",
			input:   "256.1.1.1",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip, err := NewIP(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, string(ip))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, string(ip))
			}
		})
	}
}
