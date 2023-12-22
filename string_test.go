package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	var testCases = []struct {
		name       string
		whenLength uint8
		expect     string
	}{
		{
			name:       "ok, 16",
			whenLength: 16,
		},
		{
			name:       "ok, 32",
			whenLength: 32,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			uid := String(tc.whenLength)
			assert.Len(t, uid, int(tc.whenLength))
		})
	}
}
