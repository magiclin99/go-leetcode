package p20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnswer(t *testing.T) {

	testcases := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"({)}", false},
	}

	for _, tc := range testcases {
		got := isValid(tc.input)
		assert.Equal(t, tc.expected, got)
	}
}
