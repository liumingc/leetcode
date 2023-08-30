package main

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

type BSTIterator struct {
    curr *TreeNode
    parent []*TreeNode
    visitLeft bool
}


func Constructor(root *TreeNode) BSTIterator {
    return BSTIterator{
        curr: root,
        parent: make([]*TreeNode, 0),
        visitLeft: false,
    }
}


func (this *BSTIterator) Next() int {
    curr := this.curr
    for curr.Left != nil && !this.visitLeft {
        this.parent = append(this.parent, curr)
        curr = curr.Left
    }
    res := curr.Val
    if curr.Right != nil {
        this.curr = curr.Right
        this.visitLeft = false
    } else {
        n := len(this.parent)
        if n > 0 {
            this.curr = this.parent[n-1]
            this.parent = this.parent[0:n-1]
            this.visitLeft = true
        } else {
            this.curr = nil
        }
    }
    return res
}


func (this *BSTIterator) HasNext() bool {
    return this.curr != nil
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
