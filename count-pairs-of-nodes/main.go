package main

import "fmt"

func main() {
	edges := [][]int{
		{1, 2}, {2, 4}, {1, 3}, {2, 3}, {2, 1},
	}
	res := countPairs(4, edges, []int{2, 3})
	fmt.Println(res)
}
