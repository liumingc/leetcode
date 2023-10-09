package main

import "fmt"

func main() {
	nums1 := []int{
		2, 3, 4,
	}
	nums2 := []int{
		1,
	}
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println(res)
}
