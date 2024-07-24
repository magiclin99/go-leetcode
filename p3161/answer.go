package p3161

/**
Not yet completed
*/

import "math"

type Node struct {
	Root        *Node
	Index       int
	Left        *Node
	Right       *Node
	LeftHeight  int
	RightHeight int
}

func getResults(queries [][]int) []bool {
	var db *Node

	var output []bool

	for _, q := range queries {
		if q[0] == 1 {
			// insert obstacle
			insertObstacle(&db, q[1])
		} else {
			if q[2] > q[1] {
				output = append(output, false)
			} else {
				output = append(output, hasAvailableSpace(db, q[1], q[2], 0, math.MaxInt))
			}
		}
	}

	return output
}

func insertObstacle(r **Node, index int) {

	if (*r) == nil {
		*r = &Node{
			Root:        nil,
			Index:       index,
			LeftHeight:  0,
			RightHeight: 0,
		}
		return
	}

	n := r
	for {

		currentNode := *n

		if currentNode == nil {
			*n = &Node{
				Root:        *r,
				Index:       index,
				LeftHeight:  0,
				RightHeight: 0,
			}
			break
		} else {
			if currentNode.Index == index {
				return
			} else if index < currentNode.Index {
				r = &currentNode
				n = &currentNode.Left
			} else {
				r = &currentNode
				n = &(currentNode.Right)
			}
		}
	}

	reportHeight(*n)

	rebalance(*n)
}

func reportHeight(n *Node) {
	var height int
	for {
		if n.Root == nil {
			return
		}

		height = max(n.LeftHeight, n.RightHeight) + 1

		if n.Root.Right == n {
			n.Root.RightHeight = height
		} else {
			n.Root.LeftHeight = height
		}

		n = n.Root
	}
}

func rebalance(n *Node) {
	for {
		heightDiff := (*n).LeftHeight - (*n).RightHeight
		if heightDiff > 1 {
			rightRotate(&n)
		} else if heightDiff < -1 {
			// right heavy
			leftRotate(&n)
		}

		if (*n).Root == nil {
			return
		}

		n = n.Root
	}
}

func leftRotate(r **Node) {

	ancient := (*r).Root
	current := *r
	*r = current.Right
	(*r).Root = ancient
	current.Right = nil
	current.RightHeight = 0

	leftHeight := 0
	left := &(*r).Left
	for {
		leftHeight++
		if (*left) == nil {
			*left = current
			(*r).LeftHeight = leftHeight
			return
		} else {
			left = &(*left).Left
		}
	}
}

func rightRotate(r **Node) {

	ancient := (*r).Root
	current := *r
	*r = current.Left
	(*r).Root = ancient
	current.Left = nil
	current.LeftHeight = 0

	rightHeight := 0
	right := &(*r).Right
	for {
		rightHeight++
		if (*right) == nil {
			*right = current
			(*r).RightHeight = rightHeight
			return
		} else {
			right = &(*right).Right
		}
	}
}

func hasAvailableSpace(n *Node, idxRange, reqSize, leftBound, rightBound int) bool {
	if n == nil {
		var space = rightBound - leftBound
		if idxRange < rightBound {
			space = idxRange - leftBound
		}
		return space >= reqSize
	} else {
		if n.Index >= reqSize {
			if hasAvailableSpace(n.Left, idxRange, reqSize, leftBound, n.Index) {
				return true
			}
		}

		if n.Index < idxRange {
			return hasAvailableSpace(n.Right, idxRange, reqSize, n.Index, rightBound)
		}

		return false
	}
}
