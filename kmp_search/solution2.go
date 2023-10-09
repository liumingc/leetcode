package main

import "fmt"

func search(s string, needle string) int {
	q := 0
	qn := len(needle)

	info := preCompute(needle)
	for p := 0; p < len(s); {
		debugs(s, p)
		debugs(needle, q)
		fmt.Println("----------")
		// fmt.Printf("p=%d[%c], q=%d[%c]\n", p, s[p], q, needle[q])

		if s[p] == needle[q] {
			p++
			q++
			if q >= qn {
				// found first match
				return p - qn
			}
			continue
		}
		// if not match
		if q == 0 {
			p++
			continue
		}

		nextq := info[q]
		q = nextq
	}
	return -1
}

func preCompute(needle string) []int {
	nN := len(needle)
	res := make([]int, nN)

	for p := 1; p < nN; p++ {
		for i := 0; i < nN && p+i < nN; i++ {
			if needle[i] != needle[p+i] {
				if i > res[p+i] {
					res[p+i] = i
				}
				break
			}
		}
	}

	return res
}
