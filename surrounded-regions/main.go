package main

import (
	"fmt"
)

func main() {
	/*
		board := [][]byte{
			{'X', 'X', 'X', 'X', 'O', 'X'},
			{'O', 'X', 'X', 'O', 'O', 'X'},
			{'X', 'O', 'X', 'O', 'O', 'O'},
			{'X', 'O', 'O', 'O', 'X', 'O'},
			{'O', 'O', 'X', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'O', 'X', 'X'},
		}
	*/
	board := [][]byte{
		{'O', 'O', 'O', 'O', 'X', 'X'},
		{'O', 'O', 'O', 'O', 'O', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'O'},
		{'O', 'X', 'O', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'O'},
		{'O', 'X', 'O', 'O', 'O', 'O'},
	}
	solve(board)
}

func printBoard(b [][]byte) {
	fmt.Println("____________")
	for i := 0; i < len(b); i++ {
		fmt.Printf("%s\n", b[i])
	}
	fmt.Println("____________")
}

func printVisit(b [][]bool) {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			if b[i][j] {
				fmt.Printf("x")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("")
	}
	fmt.Println("____________")
}
