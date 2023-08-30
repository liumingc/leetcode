package main

func isValidBST(root *TreeNode) bool {
    parents := make([]*TreeNode, 0)
    curr := root
    prevVal := math.MinInt
    for curr != nil || len(parents) > 0 {
        for curr != nil {
            parents = append(parents, curr)
            curr = curr.Left
        }
        pN := len(parents)
        curr, parents = parents[pN-1], parents[0:pN-1] // pop
        if curr.Val <= prevVal {
            return false
        }
        prevVal = curr.Val
        curr = curr.Right
    }
    return true
}
