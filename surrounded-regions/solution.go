package main

import "fmt"

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	bitmap := make([][]bool, m)
	for i := 0; i < m; i++ {
		bitmap[i] = make([]bool, n)
	}

	var onBoarder bool
	var dfsCheck func(y, x int)
	var dfsMark func(y, x int)

	dfsCheck = func(y, x int) {
		if board[y][x] == 'X' {
			return
		}
		if y <= 0 || y >= m-1 || x <= 0 || x >= n-1 {
			onBoarder = true
			return
		}
		if bitmap[y][x] {
			return
		}
		bitmap[y][x] = true
		dfsCheck(y, x+1)
		dfsCheck(y, x-1)
		dfsCheck(y+1, x)
		dfsCheck(y-1, x)
	}
	dfsMark = func(y, x int) {
		if y <= 0 || y >= m-1 || x <= 0 || x >= n-1 {
			return
		}
		if board[y][x] == 'X' {
			return
		}
		board[y][x] = 'X'
		dfsMark(y, x+1)
		dfsMark(y, x-1)
		dfsMark(y+1, x)
		dfsMark(y-1, x)
	}

	// iter board
	printBoard(board)
	for y := 1; y < m-1; y++ {
		for x := 1; x < n-1; x++ {
			if board[y][x] == 'O' && !bitmap[y][x] {
				onBoarder = false
				dfsCheck(y, x)
				printVisit(bitmap)
				if !onBoarder {
					fmt.Printf("(%d,%d), dfsMark\n", y, x)
					dfsMark(y, x)
					printBoard(board)
				}
			}
		}
	}
}
