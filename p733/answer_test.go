package p733

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnswer(t *testing.T) {

	testCases := []struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
		expected [][]int
	}{
		{
			image: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
			sr:       1,
			sc:       1,
			newColor: 2,
			expected: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
		{
			image: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			sr:       0,
			sc:       0,
			newColor: 0,
			expected: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			got := floodFill(tc.image, tc.sr, tc.sc, tc.newColor)
			assert.Equal(t, tc.expected, got)
		})
	}

}
