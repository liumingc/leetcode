package main

import (
	"fmt"
)

func main() {
	res := trap([]int{
		// 0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1,
		// 0, 10, 0, 70, 10, 40, 20, 30, 20, 50, 20, 40,
		// 4, 2, 0, 3, 2, 5,
		5, 5, 1, 7, 1, 1, 5, 2, 7, 6, // => 23
	})
	fmt.Println(res)
}
