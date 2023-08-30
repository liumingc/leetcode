package main

import (
	"fmt"
)

func main() {
	grid := [][]int{
		{-37,51,-36,34,-22},
		{82,4,30,14,38},
		{-68,-52,-92,65,-85},
		{-49,-3,-77,8,-19},
		{-60,-71,-21,-62,-73},
	}
	showGrid(grid)
	res := minFallingPathSum(grid)
	fmt.Println("minSum=", res)
}

func showGrid(grid [][]int) {
	n := len(grid)
	for y:=0; y<n; y++ {
		for x := 0; x < n; x++ {
			fmt.Printf("|%5d", grid[y][x])
		}
		fmt.Println("|")
	}
}
