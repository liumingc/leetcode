package main

func minCapability4(nums []int, k int) int {
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

	// O(n*k)
	for j := 1; j < k; j++ {
		// init mat[2*j][j]
		j2 := j * 2
		currSel[j2] = minArr[j2]
		for i := j2 + 1; i < n; i++ {
			currSel[i] = min(
				max(prevSel[i-2], nums[i]),
				currSel[i-1],
			)
		}
		prevSel, currSel = currSel, prevSel
	}

	return prevSel[n-1]
}
