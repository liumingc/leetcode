package main

import "fmt"

func main() {
	nums := []int{
		-4, -7, -9, -8, 6, 7, -10, -6, 9, 2,
	}
	res := maxSubarraySumCircular(nums)
	fmt.Println(res)
}
