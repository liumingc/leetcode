package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


/*
func isValidBST(root *TreeNode) bool {
	_, _, valid := validBst(root)
	return valid
}
*/

func validBst(root *TreeNode) (min, max int, valid bool) {
	min, max = root.Val, root.Val
	if root.Left != nil {
		lmin, lmax, lvalid := validBst(root.Left)
		if !lvalid {
			return
		}
		if lmax >= root.Val {
			return
		}
		min = lmin
	}
	if root.Right != nil {
		rmin, rmax, rvalid := validBst(root.Right)
		if !rvalid {
			return
		}
		if root.Val >= rmin {
			return
		}
		max = rmax
	}
	valid = true
	return
}
