package main

import "math"

func minTrioDegree(n int, edges [][]int) int {
	// adjMap := make([]map[int]bool, n+1)
	adjMap := make([][]bool, n+1)
	adjLst := make([][]int, n+1)
	minDeg := math.MaxInt

	for i := 1; i <= n; i++ {
		// adjMap[i] = make(map[int]bool)
		adjMap[i] = make([]bool, n+1)
	}

	// there are no duplicated edges in the graph
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		adjMap[from][to] = true
		adjMap[to][from] = true
		adjLst[from] = append(adjLst[from], to)
		adjLst[to] = append(adjLst[to], from)
	}
	// visited := map[trio]bool{}
	visited := make([]bool, n+1) // map[int]bool{}
	// find trios
	for x, adjs := range adjLst {
		visited[x] = true
		adjN := len(adjs)
		if adjN < 2 {
			continue
		}
		arr := adjs

		for i := 0; i < adjN-1; i++ {
			if visited[arr[i]] {
				continue
			}
			for j := i + 1; j < adjN; j++ {
				if visited[arr[j]] || !adjMap[arr[i]][arr[j]] {
					continue
				}
				a, b, c := x, arr[i], arr[j]

				deg := len(adjLst[a]) + len(adjLst[b]) + len(adjLst[c]) - 6
				if deg < minDeg {
					minDeg = deg
					if minDeg == 0 {
						return 0
					}
				}
			}
		}
	}
	if minDeg == math.MaxInt {
		return -1
	}
	return minDeg
}
