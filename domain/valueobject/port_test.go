package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPort(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{
			name:    "valid port",
			input:   22,
			want:    22,
			wantErr: false,
		},
		{
			name:    "minimum port",
			input:   1,
			want:    1,
			wantErr: false,
		},
		{
			name:    "maximum port",
			input:   65535,
			want:    65535,
			wantErr: false,
		},
		{
			name:    "port zero",
			input:   0,
			want:    0,
			wantErr: true,
		},
		{
			name:    "negative port",
			input:   -1,
			want:    0,
			wantErr: true,
		},
		{
			name:    "port too high",
			input:   70000,
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			port, err := NewPort(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, Port(0), port)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, int(port))
			}
		})
	}
}
