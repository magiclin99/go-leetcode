package p704

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {

	testCases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		}, {
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			got := search(tc.nums, tc.target)
			assert.Equal(t, tc.expected, got)
		})

	}
}
