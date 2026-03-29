package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHostName_Type(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "valid hostname",
			input: "example.com",
			want:  "example.com",
		},
		{
			name:  "empty hostname",
			input: "",
			want:  "",
		},
		{
			name:  "hostname with subdomain",
			input: "sub.example.com",
			want:  "sub.example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hn := HostName(tt.input)
			assert.Equal(t, tt.want, string(hn))
		})
	}
}
