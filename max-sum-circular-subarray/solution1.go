package main

import "fmt"

func maxSubarraySumCircular(nums []int) int {
	numL := len(nums)
	if numL == 1 {
		return nums[0]
	}

	var help func(l, r int) *sumInfo
	help = func(l, r int) *sumInfo {
		if l == r {
			return &sumInfo{
				nums[l], nums[l], nums[l], nums[l], l,
			}
		}
		mid := (l + r) / 2
		ls := help(l, mid)
		rs := help(mid+1, r)
		rsum := rs.rSum
		ri := rs.ri
		if nxtR := rs.accSum + ls.rSum; nxtR > rsum {
			rsum = nxtR
			ri = ls.ri
		}
		combine := &sumInfo{
			max(ls.lSum, ls.accSum+rs.lSum),
			rsum,
			ls.maxSum,
			ls.accSum + rs.accSum,
			ri,
		}
		combine.maxSum = max(ls.maxSum, max(
			rs.maxSum, ls.rSum+rs.lSum),
		)
		return combine
	}

	// find min neg
	currSeg := &segInfo{0, 0}
	minSeg := &segInfo{0, 0}
	minSum := nums[0]
	pre := 0
	for i := 0; i < numL; i++ {
		nxt := pre + nums[i]
		if nxt < nums[i] {
			pre = nxt
			currSeg.q = i
		} else {
			currSeg.p = i
			currSeg.q = i
			pre = nums[i]
		}
		if pre < minSum {
			minSum = pre
			minSeg.p = currSeg.p
			minSeg.q = currSeg.q
		}
	}
	fmt.Printf("minSum=%d, minSeg=%v\n", minSum, minSeg)

	res := help(0, numL-1)
	if minSum < 0 {
		l := help(0, minSeg.p)
		r := help(minSeg.q, numL-1)
		rewind := l.lSum + r.rSum
		if rewind > res.maxSum {
			return rewind
		}
	}

	return res.maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type sumInfo struct {
	lSum, rSum, maxSum, accSum int
	ri                         int // rSum start index
}

type segInfo struct {
	p, q int // index
	// sum int
}
