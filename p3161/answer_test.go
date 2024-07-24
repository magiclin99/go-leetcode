package p3161

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetResults(t *testing.T) {

	testCases := []struct {
		name     string
		queries  [][]int
		expected []bool
	}{
		{
			name: "Test Case 1",
			queries: [][]int{
				{1, 2},
				{2, 3, 3},
				{2, 3, 1},
				{2, 2, 2},
			},
			expected: []bool{false, true, true},
		},
		{
			name: "Test Case 2",
			queries: [][]int{
				{1, 7},
				{2, 7, 6},
				{1, 2},
				{2, 7, 5},
				{2, 1, 2},
			},
			expected: []bool{true, true, false},
		},
		{
			name: "Test Case 3",
			queries: [][]int{
				{2, 1, 2},
			},
			expected: []bool{false},
		},
		{
			name: "Test Case 4",
			queries: [][]int{
				{1, 3},
				{2, 4, 2},
			},
			expected: []bool{true},
		},
		{
			name: "Test Case 5",
			queries: [][]int{
				{1, 1},
				{1, 5},
				{1, 13},
				{1, 14},
				{2, 12, 8},
			},
			expected: []bool{false},
		},
		{
			name: "Test Case 6",
			queries: [][]int{
				{1, 1},
				{1, 11},
				{1, 4},
				{1, 8},
				{2, 13, 7},
			},
			expected: []bool{false},
		},
		{
			name: "Test Case 7",
			queries: [][]int{
				{1, 7},
				{1, 1},
				{1, 3},
				{2, 3, 12},
				{2, 3, 3},
			},
			expected: []bool{false, false},
		},
		{
			name: "Test Case 8",
			queries: [][]int{
				{1, 12},
				{1, 2},
				{2, 9, 7},
				{1, 3},
				{2, 6, 6},
			},
			expected: []bool{true, false},
		},
		{
			name: "Test Case 9",
			queries: [][]int{
				{1, 1},
				{1, 2},
				{1, 7},
				{1, 5},
				{1, 3},
				{1, 8},
				{2, 5, 2},
			},
			expected: []bool{true},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getResults(tc.queries)
			assert.Equal(t, tc.expected, result)
		})
	}
}
