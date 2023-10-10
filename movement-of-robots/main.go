package main

import "fmt"

func main() {
	nums := []int{-2, 0, 2}
	s := "RLL"
	d := 3
	res := sumDistance(nums, s, d)
	fmt.Println("res is ", res)
}
