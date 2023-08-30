package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
    res, _ := kth(root, 1, k)
    if res != nil {
        return res.Val
    }
    return 0
}

func kth(root *TreeNode, i, k int) (res *TreeNode, nextI int) {
    nextI = i
    if root.Left != nil {
        res, nextI = kth(root.Left, i, k)
        if res != nil {
            return res, nextI
        }
    }
    if nextI == k {
        return root, nextI
    }
    if root.Right == nil {
        return nil, nextI+1
    }
    return kth(root.Right, nextI+1, k)
}
