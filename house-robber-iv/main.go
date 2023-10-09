package main

import "fmt"

func main() {
	nums := []int{
		// 7, 3, 9, 5,
		2, 7, 9, 3, 1,
		//4, 22, 11, 14, 25,
		// 24, 1, 55, 46, 4, 61, 21, 52,
	}
	k := 2
	fmt.Println("nums=", nums, ", k=", k)
	res := minCapability(nums, k)
	fmt.Println(res)
}
