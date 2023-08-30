package main

/*
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	switch n {
	case 0:
		return nil
	case 1:
		return &TreeNode{Val: preorder[0]}
	}

	tbl := make(map[int]int, n)
	for i, x := range inorder {
		tbl[x] = i
	}
	root := &TreeNode{Val: preorder[0]}
	res := root
	rootI := tbl[root.Val]

	for i := 1; i < n; i++ {
		p := tbl[preorder[i]]
		curr := root
		xN := &TreeNode{Val: preorder[i]}
		// link the 2 nodes
		for {
			q := tbl[curr.Val]
			if p < q {
				// go left
				if curr.Left == nil {
					curr.Left = xN
					break
				}
				curr = curr.Left
				continue
			}
			// go right
			if curr.Right == nil {
				curr.Right = xN
				break
			}
			curr = curr.Right
		}
		if p > rootI {
			root = root.Right
			rootI = tbl[root.Val]
		}
	}
	return res
}
*/
