package main

import (
	"fmt"
)

func main() {
	/*
	grid := [][]byte{
		{'1', '0', '1', '1', '0', '0', '1', '0', '1', '1', '1', '1', '0', '1', '0', '1', '1', '1', '1', '0'},
		{'0', '1', '0', '0', '1', '0', '1', '0', '1', '1', '1', '1', '1', '1', '0', '1', '1', '0', '1', '1'},
		{'1', '0', '0', '1', '0', '1', '0', '1', '0', '1', '1', '0', '1', '1', '1', '0', '0', '1', '1', '0'},
		{'0', '1', '1', '0', '0', '1', '1', '0', '1', '1', '1', '1', '0', '0', '1', '0', '0', '0', '1', '1'},
		{'1', '1', '0', '1', '0', '0', '1', '0', '0', '0', '1', '0', '1', '0', '1', '1', '1', '0', '1', '1'},
		{'0', '0', '0', '0', '1', '0', '1', '1', '0', '0', '1', '0', '0', '1', '0', '1', '1', '1', '1', '0'},
		{'1', '0', '1', '1', '1', '1', '0', '1', '1', '0', '1', '1', '0', '1', '1', '1', '0', '0', '1', '0'},
		{'0', '1', '1', '0', '0', '0', '1', '0', '0', '1', '0', '1', '1', '1', '0', '0', '1', '1', '0', '1'},
		{'0', '0', '0', '0', '1', '1', '0', '1', '0', '0', '1', '1', '0', '1', '0', '0', '1', '0', '1', '0'},
		{'0', '0', '1', '1', '1', '0', '1', '0', '1', '0', '1', '1', '1', '0', '1', '1', '1', '1', '1', '0'},
		{'1', '0', '1', '0', '1', '1', '1', '0', '1', '1', '1', '0', '1', '0', '1', '0', '1', '0', '1', '1'},
		{'0', '0', '1', '1', '1', '1', '0', '1', '1', '1', '0', '1', '0', '0', '0', '1', '1', '1', '0', '1'},
		{'1', '1', '1', '0', '0', '0', '0', '0', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1', '1', '0'},
		{'0', '0', '1', '1', '1', '0', '0', '1', '0', '0', '1', '1', '1', '1', '1', '1', '0', '1', '1', '0'},
		{'0', '0', '0', '1', '1', '0', '0', '0', '0', '1', '1', '0', '1', '0', '0', '1', '1', '1', '1', '1'},
		{'0', '1', '1', '1', '0', '1', '0', '0', '1', '1', '1', '1', '1', '0', '1', '1', '1', '0', '0', '1'},
		{'0', '0', '0', '0', '1', '1', '1', '1', '0', '0', '0', '0', '1', '0', '0', '0', '0', '1', '1', '0'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '0', '1', '1', '1', '1', '1', '1'},
		{'0', '1', '0', '0', '1', '0', '0', '1', '1', '1', '1', '1', '1', '0', '1', '0', '1', '1', '1', '1'},
		{'0', '0', '1', '1', '1', '1', '1', '0', '0', '0', '1', '1', '1', '1', '1', '1', '0', '1', '1', '0'},
	}
	*/
	grid := [][]byte{
		{'1','0','1','1','1'},
		{'1','0','1','0','1'},
		{'1','1','1','0','1'},
	}
	res := numIslands(grid)
	fmt.Println(res)
}

func printMark(mark [][]int) {
	for y:=1; y<len(mark); y++ {
		for x:=1; x<len(mark[0]); x++ {
			if mark[y][x] == 0 {
				fmt.Printf("|..")
			} else {
				fmt.Printf("|%02d", mark[y][x])
			}
		}
		fmt.Println("")
	}
}
