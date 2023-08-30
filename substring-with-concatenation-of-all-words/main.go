package main

import "fmt"

func main() {
	/*
		s := "bcabbcaabbccacacbabccacaababcbb"
		words := []string{
			"c", "b", "a", "c", "a", "a", "a", "b", "c",
			//"fooo", "barr", "wing", "ding", "wing",
			// "word", "good", "best", "good",
		}
	*/
	s := "ababababab"
	words := []string{
		"ababa", "babab",
	}

	fmt.Println("s=", s, ", words: ", words)
	res := findSubstring(s, words)
	for _, x := range res {
		fmt.Println("=> ", x)
	}
}
