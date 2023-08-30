package main

import "fmt"

func main() {
	s := "abcde"
	idxs := []int{2, 2}
	sources := []string{"cdef", "bc"}
	targets := []string{"f", "fe"}
	res := findReplaceString(s, idxs, sources, targets)
	fmt.Println(s, idxs, sources, targets, " => ", res)
}
