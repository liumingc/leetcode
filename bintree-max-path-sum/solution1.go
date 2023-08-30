package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    maxSum = -1001
    recurSum(root)
    return maxSum
}

var maxSum int

func recurSum(root *TreeNode) (res int) {
    if root == nil {
        return 0
    }
    res = root.Val
    left := recurSum(root.Left)
    right := recurSum(root.Right)
    switch {
    case left <= 0 && right <= 0:
        // do nothing
    case left <= 0:
        // right > 0
        res += right
    case right <= 0:
        res += left
    default:
        // left > 0 && right > 0
        tmp := left + res + right
        if tmp > maxSum {
            maxSum = tmp
        }
        if left > right {
            res += left
        } else {
            res += right
        }
    }
    if res > maxSum {
        maxSum = res
    }
    return res
}
