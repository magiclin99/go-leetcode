package p235

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack []*TreeNode

func (it *Stack) Push(value *TreeNode) {
	*it = append(*it, value)
}

func (it *Stack) Pop() {
	*it = (*it)[:len(*it)-1]
}

func (it *Stack) Top() *TreeNode {
	last := len(*it) - 1
	if last >= 0 {
		return (*it)[last]
	}
	return nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	targets := map[int]bool{
		p.Val: true,
		q.Val: true,
	}
	ans, _ := eliminateTarget(root, targets)
	return ans
}

func eliminateTarget(root *TreeNode, targets map[int]bool) (ans *TreeNode, score int) {
	if root == nil {
		return nil, 0
	}

	totalScore := 0

	ans, got := eliminateTarget(root.Left, targets)
	if ans != nil {
		return ans, -1
	}
	totalScore += got

	ans, got = eliminateTarget(root.Right, targets)
	if ans != nil {
		return ans, -1
	}
	totalScore += got

	// now, score can be 0, 1, 2
	if totalScore == 2 {
		return root, -1
	}

	// now, score can be 0, 1
	_, ok := targets[root.Val]
	if ok {
		delete(targets, root.Val)
		totalScore++
	}

	// now score can be 1 or 2
	if totalScore == 2 {
		return root, -1
	}

	// now score can be 1
	return nil, totalScore
}
