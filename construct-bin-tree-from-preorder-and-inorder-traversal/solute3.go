package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	switch n {
	case 0:
		return nil
	case 1:
		return &TreeNode{Val: preorder[0]}
	}

	a := preorder[0]
	pre := &TreeNode{Val: a}
	res := pre
	stack := []*TreeNode{pre}
	pN := 1 // pre order index
	iN := 0 // in order index
	for pN < n && iN < n {
		b := inorder[iN]
		sN := len(stack) - 1
		if sN >= 0 && b == stack[sN].Val {
			for sN > 0 && iN < n-1 &&
				inorder[iN+1] == stack[sN-1].Val {
				iN++
				sN--
			}
			pre = stack[sN]
			stack = stack[0:sN]
			iN++
			// go right
			xN := &TreeNode{Val: preorder[pN]}
			pre.Right = xN
			pre = xN
			pN++
			stack = append(stack, xN)
			continue
		}
		a := preorder[pN]
		xN := &TreeNode{Val: a}
		// link left
		pre.Left = xN
		pre = xN
		stack = append(stack, xN)
		pN++
	}
	return res
}
