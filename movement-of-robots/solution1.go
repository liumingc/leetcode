package main

import (
	"fmt"
	"sort"
)

func sumDistance(nums []int, s string, d int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		switch s[i] {
		case 'L':
			nums[i] -= d
		case 'R':
			nums[i] += d
		}
	}
	res := 0
	/*
	   for i:=0; i<n-1; i++ {
	       for j:=i+1; j<n; j++ {
	           res += abs(nums[i] - nums[j])
	       }
	   }
	*/
	sort.Ints(nums)
	fmt.Println("nums = ", nums)
	curr := nums[1] - nums[0]
	fmt.Printf("i=1, curr=%d\n", curr)
	res += abs(curr)
	accN := 1
	for i := 2; i < n; i++ {
		accN++
		diff := nums[i] - nums[i-1]
		curr = curr + accN*diff
		res = (res + abs(curr)) % dround
		fmt.Printf("i=%d, diff=%d, curr=%d, res=%d\n", i, diff, curr, res)
	}
	return res % dround
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

const (
	dround = 1e9 + 7
)
