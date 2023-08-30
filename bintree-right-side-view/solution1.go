package main

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    res := make([]int, 0)
    arr := []*TreeNode{root}
    for len(arr) > 0 {
        nextArr := make([]*TreeNode, 0)
        for _, x := range arr {
            if x.Right != nil {
                nextArr = append(nextArr, x.Right)
            }
            if x.Left != nil {
                nextArr = append(nextArr, x.Left)
            }
        }
        res = append(res, arr[0].Val) // only get the right most one
        arr = nextArr
    }
    return res
}
