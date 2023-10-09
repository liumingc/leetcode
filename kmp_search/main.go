package main

import "fmt"

func main() {
	arr := []string{
		"ababc",
		"aaaa",
		"abababcdab",
		"abcde",
		"issip",
	}
	for _, x := range arr {
		preRes := preCompute(x)
		fmt.Println(x, " => ", preRes)
		for i, pr := range preRes {
			if pr == 0 {
				continue
			}
			shows(x, i, pr)
		}
	}
	fmt.Println("end of preCompute")

	inputArr := []string{
		"aaaababcde",
		"aaabbbaaaabb",
		"abc",
		"abcdgabcdfabcdem",
		"mississippi",
	}
	for i, input := range inputArr {
		n := search(input, arr[i])
		if n >= 0 {
			fmt.Printf("[%d] pat=%s, res=%s\n", i, arr[i], input[n:])
		}
	}
}

func debugs(s string, k int) {
	fmt.Println(s)
	for i := 0; i < k; i++ {
		fmt.Printf(" ")
	}
	fmt.Println("^")
}

func shows(s string, p, k int) {
	fmt.Printf("shows, s=%s, p=%d, k=%d\n", s, p, k)
	if k <= 0 {
		return
	}

	for i := 0; i < k; i++ {
		fmt.Printf(" ")
	}
	fmt.Println(s)
	fmt.Println(s)
	for i := 0; i < p; i++ {
		fmt.Printf(" ")
	}
	fmt.Println("^")
}
