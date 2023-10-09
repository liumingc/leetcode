package main

import "fmt"

func minCapability3(nums []int, k int) int {
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

	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, k)
	}

	mat[0][0] = nums[0]
	for i := 1; i < n; i++ {
		mat[i][0] = min(mat[i-1][0], nums[i])
	}
	// mat[0][0..k]
	// mat[1][0..k]
	minArr := make([]int, n)
	minArr[0] = nums[0]
	for i := 2; i < n; i += 2 {
		minArr[i] = max(minArr[i-2], nums[i])
	}

	for j := 1; j < k; j++ {
		// init mat[2*j][j]
		j2 := j * 2
		mat[j2][j] = minArr[j2]
		for i := j2 + 1; i < n; i++ {
			mat[i][j] = min(
				max(mat[i-2][j-1], nums[i]),
				mat[i-1][j],
			)
		}
	}

	fmt.Println(nums, ", ", k)
	printMat(mat)

	return mat[n-1][k-1]
}

func printMat(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			fmt.Printf("% 5d, ", mat[i][j])
		}
		fmt.Println()
	}
}
