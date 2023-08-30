package main

import "fmt"

func maxSizeSlices(slices []int) int {
	return max(sumMax(slices[0:len(slices)-1]),
		sumMax(slices[1:]))
}

func sumMax(s []int) int {
	sn := len(s)
	n := sn + 1
	tri := n / 3

	mat := make([][]int, sn)
	for i := 0; i < sn; i++ {
		mat[i] = make([]int, tri)
	}
	for j := 0; j < tri; j++ {
		mat[0][j] = s[0]
		mat[1][j] = max(s[0], s[1])
	}
	for i := 1; i < sn; i++ {
		mat[i][0] = max(mat[i-1][0], s[i])
	}

	for i := 2; i < sn; i++ {
		for j := 1; j < tri; j++ {
			mat[i][j] = max(mat[i-2][j-1]+s[i],
				mat[i-1][j])
		}
	}
	fmt.Println("slice=", s)
	printMat(mat)
	return mat[sn-1][tri-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printMat(mat [][]int) {
	yN := len(mat)
	xN := len(mat[0])
	for y := 0; y < yN; y++ {
		for x := 0; x < xN; x++ {
			fmt.Printf("| %5d", mat[y][x])
		}
		fmt.Println("")
	}
}
