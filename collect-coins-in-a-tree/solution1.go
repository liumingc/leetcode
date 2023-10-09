package main

import "fmt"

func collectTheCoins1(coins []int, edges [][]int) int {
	n := len(coins)
	visited := make([]bool, n)
	routeMap := make([][]int, n)
	for _, e := range edges {
		f, t := e[0], e[1]
		routeMap[f] = append(routeMap[f], t)
		routeMap[t] = append(routeMap[t], f)
	}
	start := -1
	for i, c := range coins {
		if c == 1 {
			start = i
			break
		}
	}
	if start == -1 {
		return 0 // no coins
	}
	res := 0
	var dfs func(i, d int) int
	br0 := 0
	b0 := 0
	br1 := 0
	found0, found1 := false, false
	dfs = func(i, d int) int {
		if visited[i] {
			return 0
		}
		// fmt.Printf("[visit coin %d, d=%d\n", i, d)
		visited[i] = true
		dist := d
		for _, nxt := range routeMap[i] {
			if visited[nxt] {
				continue
			}
			// visited[nxt] = true
			distNxt := dfs(nxt, d+1)
			if distNxt > dist {
				dist = distNxt
			}
			nxtD := distNxt - d
			if nxtD > 2 {
				fmt.Printf("\tcoin %d ->  %d, add 2\n", i, nxt)
				res += 2 // forward and backward
			}
			switch d {
			case 0:
				if nxtD > 0 {
					b0++
				}
				fmt.Printf("\t .. coin %d -> %d, nxtD=%d\n", i, nxt, nxtD)
				if nxtD >= 2 {
					br0++
				}
			case 1:
				if nxtD >= 2 {
					br1++
				}
			}
		}
		switch d {
		case 0:
			if br0 == 1 {
				found0 = true
			}
		case 1:
			if br1 == 1 {
				found1 = true
			}
			br1 = 0
		}
		if d == dist && coins[i] == 0 {
			return 0
		}
		// fmt.Printf("coin %d, dist=%d]\n", i, dist)
		return dist
	}
	dfs(start, 0)
	subt := 0
	if found0 {
		subt += 2
		if found1 && b0 == 1 {
			subt += 2
		}
	}
	fmt.Printf("res=%d, subt=%d\n", res, subt)
	if res >= subt {
		return res - subt
	}
	if subt == 4 && res == 2 {
		return 0
	}
	return res
}
