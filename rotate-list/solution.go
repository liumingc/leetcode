package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	n := 0
	var pre, bn *backNode
	for p := head; p != nil; p = p.Next {
		bn = &backNode{pre: pre, cur: p}
		pre = bn
		n++
	}
	k = k % n
	if k == 0 {
		return head
	}
	bn.cur.Next = head // link last to first
	for i := 0; i < k; i++ {
		bn = bn.pre
	}
	res := bn.cur.Next
	bn.cur.Next = nil // cut the link
	return res
}

type backNode struct {
	pre *backNode
	cur *ListNode
}
