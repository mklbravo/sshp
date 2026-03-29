package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsername_GetValue(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "valid username",
			input: "user",
			want:  "user",
		},
		{
			name:  "empty username",
			input: "",
			want:  "",
		},
		{
			name:  "username with numbers",
			input: "user123",
			want:  "user123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			un := Username(tt.input)
			got := un.GetValue()
			assert.Equal(t, tt.want, got)
		})
	}
}
