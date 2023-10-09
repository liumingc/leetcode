package main

import "fmt"

func search2(s string, needle string) int {
	q := 0
	qn := len(needle)
	info := make([]int, qn)
	/*
		for i := 0; i < qn; i++ {
			info[i] = qn - i - 1
		}
	*/
	findQ2(needle, qn, info)
	fmt.Println("... s=", s, ", pat=", needle, "info=", info)
	for p := 0; p < len(s); {
		debugs(s, p)
		debugs(needle, q)
		fmt.Println("=====")
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

		nextq := info[q-1]
		if nextq > 0 {
			p = p - q + 1 + nextq
		} // else no need to go back
		q = nextq
	}
	return -1
}

// time: O(N^3)
func findQ2(s string, k int, res []int) {
	// s[0 .. k-1]
	if k < 2 {
		return
	}
	q := k - 1
	for ; q > 0; q-- {
		i := 0
		for ; i < q; i++ {
			// fmt.Printf("k=%d, i=%d, q=%d, k-1-q+i=%d, left=<%c>, right=<%c>\n",
			//	k, i, q, k-q+i, s[i], s[k-q+i])
			if s[i] != s[k-q+i] {
				break
			}
		}
		if i == q {
			// found
			if q > res[k-1] {
				res[k-1] = q
			}
			break
		}
	}

	k--
	if k >= 2 {
		findQ2(s, k, res)
	}
}
