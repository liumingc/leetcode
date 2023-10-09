package main

import "fmt"

func minCapability(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if k == 1 {
		res := nums[0]
		for _, x := range nums[1:] {
			if x < res {
				res = x
			}
		}
		return res
	}

	currSel := make([]int, n)
	prevSel := make([]int, n)

	prevSel[0] = nums[0]
	for i := 1; i < n; i++ {
		prevSel[i] = min(prevSel[i-1], nums[i])
	}

	minArr := make([]int, n)
	minArr[0] = nums[0]
	for i := 2; i < n; i += 2 {
		minArr[i] = max(minArr[i-2], nums[i])
	}
	lastArr := make([]int, n)
	lastArr[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		lastArr[i] = min(lastArr[i+1], nums[i])
	}
	fmt.Println("lastArr, ", lastArr)

	type maxInfo struct {
		i   int
		val int
	}
	prevMax := maxInfo{i: n}
	currMax := maxInfo{i: n}
	// O(n*k)
	for j := 1; j < k; j++ {
		// init mat[2*j][j]
		j2 := j * 2
		currSel[j2] = minArr[j2]
		currMax.i = j2
		currMax.val = currSel[j2]
		early := false
		for i := j2 + 1; i < n; i++ {
			if currSel[i-1] < lastArr[i] {
				early = true
				currMax.i = i - 1
				currMax.val = currSel[i-1]
				break
			}
			var prev int
			if i-2 < prevMax.i {
				prev = prevSel[i-2]
			} else {
				prev = prevMax.val
			}
			fmt.Printf("\ti=%d, prev=%d, nums=%d\n", i, prev, nums[i])
			currSel[i] = min(
				max(prev, nums[i]),
				currSel[i-1],
			)
		}
		if !early {
			currMax.i = n - 1
			currMax.val = currSel[n-1]
		}
		prevSel, currSel = currSel, prevSel
		prevMax = currMax
		fmt.Println(j, "currMax, ", prevMax)
	}

	return currMax.val
}
