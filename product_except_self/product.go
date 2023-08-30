package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	// front := make([]int, n)
	rear := make([]int, n)
	/*
	   front[0] = 1
	   for i:=0; i<n-1; i++ {
	       front[i+1] = front[i] * nums[i]
	       if front[i+1] == 0 {
	           break
	       }
	   }
	*/
	rear[n-1] = 1
	for i := n - 1; i > 0; i-- {
		rear[i-1] = rear[i] * nums[i]
		if rear[i-1] == 0 {
			break
		}
	}
	// res := make([]int, n)
	front := 1
	prior := 1
	for i := 0; i < n; i++ {
		front = front * prior
		res := front * rear[i]
		prior = nums[i]
		nums[i] = res
	}
	return nums
}
