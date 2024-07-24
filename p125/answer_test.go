package p125

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s      string
		expect bool
	}{
		{
			s:      "race a car",
			expect: false,
		},
		{
			s:      "A man, a plan, a canal: Panama",
			expect: true,
		},
		{
			s:      "0P",
			expect: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.s, func(t *testing.T) {
			result := isPalindrome(tc.s)
			assert.Equal(t, tc.expect, result)
		})
	}

}
