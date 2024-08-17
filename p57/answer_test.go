package p57

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"case 1",
			args{
				[][]int{{1, 3}, {6, 9}},
				[]int{2, 5},
			},
			[][]int{{1, 5}, {6, 9}},
		}, {
			"case 2",
			args{
				[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				[]int{4, 8},
			},
			[][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			"case 3",
			args{
				[][]int{},
				[]int{5, 7},
			},
			[][]int{{5, 7}},
		},
		{
			"case 4",
			args{
				[][]int{{1, 5}},
				[]int{6, 8},
			},
			[][]int{{1, 5}, {6, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outptu := insert(tt.args.intervals, tt.args.newInterval)
			assert.Equal(t, tt.want, outptu)
		})
	}
}
