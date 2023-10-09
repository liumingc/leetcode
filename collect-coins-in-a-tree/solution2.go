package main

func collectTheCoins(coins []int, edges [][]int) int {
	n := len(coins)
	routeMap := make([][]int, n)
	for _, e := range edges {
		f, t := e[0], e[1]
		routeMap[f] = append(routeMap[f], t)
		routeMap[t] = append(routeMap[t], f)
	}
	start := -1
	coinN := 0
	for i, c := range coins {
		if c == 1 {
			coinN++
			if start == -1 {
				start = i
			}
		}
	}
	if start == -1 {
		return 0 // no coins
	}
	res := 0
	var dfs func(prev, i, d int) int
	br0 := 0
	b0 := 0
	br1 := 0
	found0, found1 := false, false
	dfs = func(prev, i, d int) int {
		dist := d
		for _, nxt := range routeMap[i] {
			if nxt == prev {
				continue
			}
			distNxt := dfs(i, nxt, d+1)
			if coinN == 0 {
				break
			}
			if distNxt > dist {
				dist = distNxt
			}
			nxtD := distNxt - d
			if nxtD > 2 {
				res += 2 // forward and backward
			}
			switch d {
			case 0:
				if nxtD > 0 {
					b0++
				}
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
		if coins[i] == 1 {
			coinN--
		}
		return dist
	}
	dfs(-1, start, 0)
	subt := 0
	if found0 {
		subt += 2
		if found1 && b0 == 1 {
			subt += 2
		}
	}
	if res >= subt {
		return res - subt
	}
	if res == 2 && subt == 4 {
		return 0
	}
	return res
}
