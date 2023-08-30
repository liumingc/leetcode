package main

import (
	"fmt"
)

func candy(ratings []int) int {
	n := len(ratings)
	if n <= 1 {
		return 1
	}
	tot := 0
	curr := 1

	for i := 0; i < n; {
		p := i
		for ; p < n-1; p++ {
			if ratings[p] > ratings[p+1] {
				continue
			} else {
				break
			}
		}
		step := p - i
		if step > 0 {
			peak := 1 + step
			adjust := 0
			if i > 0 && ratings[i] > ratings[i-1] {
				curr += 1
				if curr > peak {
					adjust = curr - peak
					fmt.Printf("[%d] adjust %d\n", i, adjust)
				}
			}

			curr = peak
			for ; i < p; i++ {
				tot += curr
				fmt.Printf("[%d] %d -> %d\n", i, curr, tot)
				curr--
			}
			tot += adjust
		}
		if i > 0 && ratings[i] > ratings[i-1] {
			curr += 1
		} else {
			curr = 1 // reset
		}
		tot += curr
		fmt.Printf("[%d] %d -> %d\n", i, curr, tot)
		i++
	}
	return tot
}
