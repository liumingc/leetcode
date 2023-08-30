package main

import "fmt"

func trap(height []int) int {
	n := len(height)
	if n <= 1 {
		return 0
	}

	// collect top infos
	topi := 0
	tops := make([]*Top, n)
	for i := 0; i < n; i++ {
		if isTop(height, i) {
			if topi > 0 {
				// go backword
				ii := topi - 1
				for ; ii > 0; ii-- {
					if tops[ii].ht >= height[i] {
						break
					}
					if tops[ii].ht > tops[ii-1].ht {
						break
					}
				}
				topi = ii + 1
			}
			tops[topi] = &Top{
				i:  i,
				ht: height[i],
			}
			topi++
		}

	}
	for i := 0; i < topi; i++ {
		fmt.Println(tops[i].i, " => ", tops[i].ht)
	}

	vol := 0
	// acc from top to top
	for i := 0; i < topi-1; i++ {
		leftTop := tops[i]
		rightTop := tops[i+1]
		lev := leftTop.ht
		if rightTop.ht < lev {
			lev = rightTop.ht
		}
		fmt.Printf("left=%d, right=%d, lev=%d\n", leftTop.ht, rightTop.ht, lev)
		for ii := leftTop.i; ii < rightTop.i; ii++ {
			x := lev - height[ii]
			if x > 0 {
				vol += x
			}
			fmt.Printf("[%d] => %d\n", ii, x)
		}
	}
	return vol
}

type Top struct {
	i  int
	ht int
}

func isTop(hts []int, i int) bool {
	n := len(hts)
	if i == 0 {
		if hts[0] >= hts[1] {
			return true
		}
	} else if i == n-1 {
		if hts[n-2] < hts[n-1] {
			return true
		}
	} else {
		if hts[i-1] < hts[i] && hts[i] >= hts[i+1] {
			return true
		}
	}
	return false
}
