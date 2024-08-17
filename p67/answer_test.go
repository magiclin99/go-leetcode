package p67

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBinary(t *testing.T) {
	testCases := []struct {
		a        string
		b        string
		expected string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s+%s", tc.a, tc.b), func(t *testing.T) {
			actual := addBinary(tc.a, tc.b)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
