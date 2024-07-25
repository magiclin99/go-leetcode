package p235

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnswer(t *testing.T) {

	testCases := []struct {
		tree     []int
		p        int
		q        int
		expected int
	}{
		{
			tree:     []int{6, 2, 8, 0, 4, 7, 9, 3, 5},
			p:        2,
			q:        8,
			expected: 6,
		},
		{
			tree:     []int{6, 2, 8, 0, 4, 7, 9, 3, 5},
			p:        2,
			q:        4,
			expected: 2,
		},

		{
			tree:     []int{2, 1},
			p:        2,
			q:        1,
			expected: 2,
		},
		{
			tree:     []int{3, 1, 4, 2},
			p:        2,
			q:        3,
			expected: 3,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			got := lowestCommonAncestor(buildBST(tc.tree), &TreeNode{Val: tc.p}, &TreeNode{Val: tc.q}).Val
			assert.Equal(t, tc.expected, got)
		})
	}
}

func buildBST(tree []int) *TreeNode {
	root := &TreeNode{Val: tree[0]}
	for i := 1; i < len(tree); i++ {
		insertBST(root, tree[i])
	}
	return root
}

func insertBST(root *TreeNode, i int) {
	if i < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: i}
		} else {
			insertBST(root.Left, i)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Val: i}
		} else {
			insertBST(root.Right, i)
		}
	}
}
