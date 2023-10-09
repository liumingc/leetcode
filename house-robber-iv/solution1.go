package main

import "math"

// timeout
func minCapability1(nums []int, k int) int {
	n := len(nums)
	var selek func(si, leftK int) int
	cache := make(map[key]int)
	selek = func(si, leftK int) int {
		ki := key{si, leftK}
		if res, ok := cache[ki]; ok {
			return res
		}
		res := nums[si]
		if leftK == 1 {
			for i := si + 1; i < n; i++ {
				res = min(res, nums[i])
			}
			cache[ki] = res
			return res
		}
		// is it possible to select leftK elems
		// from nums[i .. n-1]?
		res = math.MaxInt
		leftN := n - (leftK + leftK - 1)
		for i := si; i <= leftN; i++ {
			tmpMax := max(nums[i], selek(i+2, leftK-1))
			res = min(res, tmpMax)
		}
		cache[ki] = res
		return res
	}
	return selek(0, k)
}

type key struct {
	si, leftK int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
