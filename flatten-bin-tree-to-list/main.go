package main

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}
	flatten(root)
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func flatten(root *TreeNode) {
	flat(root)
}

func flat(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	leftTail := flat(root.Left)
	rightTail := flat(root.Right)
	if leftTail != nil {
		leftTail.Right = root.Right
		root.Right = root.Left
		root.Left = nil
		if rightTail == nil {
			return leftTail
		}
	}
	if rightTail != nil {
		return rightTail
	}
	return root
}
