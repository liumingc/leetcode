package main

import "container/heap"

func mergeKLists2(lists []*ListNode) *ListNode {
	nL := len(lists)
	if nL == 0 {
		return nil
	}
	if nL == 1 {
		return lists[0]
	}

	lh := &LH{}
	heap.Init(lh)
	for i := 0; i < nL; i++ {
		if lists[i] == nil {
			continue
		}
		heap.Push(lh, lists[i])
	}

	dummy := &ListNode{}
	pre := dummy

	for lh.Len() > 0 {
		sel := heap.Pop(lh).(*ListNode)
		pre.Next = sel
		pre = sel
		nxt := sel.Next

		if lh.Len() == 0 {
			break
		}

		nxtMin := (*lh)[0].Val
		for nxt != nil && nxt.Val <= nxtMin {
			pre.Next = nxt
			pre = nxt
			nxt = nxt.Next
		}
		if nxt != nil {
			heap.Push(lh, nxt)
		}
		// sel.Next = nil
	}
	return dummy.Next
}
