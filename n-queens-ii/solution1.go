package main

import "fmt"

func totalNQueens(n int) int {
	arr := make([]int, n)
	put := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	res := 0
	var backtrack func(int)
	backtrack = func(r int) {
		printBoard(arr, put, r, n)
		if r == n {
			res++
			return
		}
		for i := r; i < n; i++ {
			if i != r {
				arr[i], arr[r] = arr[r], arr[i] // swap
			}

			failTry := false
			for j := 0; j < r; j++ {
				if crossAttack(j, put[j], r, arr[r]) {
					failTry = true
					break
				}
			}
			fmt.Printf("try: i=%d, r=%d, arr=%v, put=%v, failTry=%v\n", i, r, arr, put, failTry)
			if !failTry {
				put[r] = arr[r]
				backtrack(r + 1)
			}
			if i != r {
				arr[i], arr[r] = arr[r], arr[i] // swap back
			}
		}
	}
	backtrack(0)
	return res
}

func crossAttack(r1, c1, r2, c2 int) bool {
	if c1 > c2 {
		c1, c2 = c2, c1
	}
	return r2-r1 == c2-c1
}

func printBoard(arr, put []int, r, n int) {
	fmt.Println(arr, "; ", put, "r=", r)
	// draw board
	for i := 0; i < r; i++ {
		for j := 0; j < put[i]; j++ {
			fmt.Printf("O")
		}
		fmt.Printf("#")
		for j := put[i] + 1; j < n; j++ {
			fmt.Printf("O")
		}
		fmt.Println("")
	}
	for i := r; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("O")
		}
		fmt.Println("")
	}
	fmt.Println("-------------------")
}
