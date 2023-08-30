package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

/*
func buildTree(preorder []int, inorder []int) *TreeNode {
	tbl := make(map[int]int, len(preorder))
	for i, x := range inorder {
		tbl[x] = i
	}
	var bt func(ps, is, n int) *TreeNode
	bt = func(ps, is, n int) *TreeNode {
		switch n {
		case 0:
			return nil
		case 1:
			return &TreeNode{Val: preorder[ps]}
		}
		x := preorder[ps]
		nextIs := tbl[x]
		nextN := nextIs - is
		nextPs := ps + nextN + 1
		left := bt(ps+1, is, nextN)
		right := bt(nextPs, nextIs+1, n-1-nextN)
		return &TreeNode{
			Val:   x,
			Left:  left,
			Right: right,
		}
	}
	return bt(0, 0, len(preorder))
}
*/
