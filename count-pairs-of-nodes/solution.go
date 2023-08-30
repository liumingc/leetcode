package main

import "sort"

func countPairs(n int, edges [][]int, queries []int) []int {
	hpair := make(map[int]map[int]int) // a -> b -> edges count
	neMap := make(map[int]bool)        // non empty node
	h := make([]int, n+1)
	eN := len(edges)

	for i := 0; i < eN; i++ {
		edge := edges[i]
		a, b := edge[0], edge[1]
		h[a]++
		h[b]++
		if a > b {
			a, b = b, a
		}
		if pair, ok := hpair[a]; ok {
			pair[b]++
		} else {
			hpair[a] = map[int]int{b: 1}
		}
		neMap[a] = true
		neMap[b] = true
	}

	// IF iter all pairs, O(n*(n-1)/2)

	// inciN := make([]int, eN+1)
	inciN := make(map[int]int)

	neArr := map2orderList(neMap)
	neN := len(neArr)
	xN := n - neN
	for i := 0; i < neN; i++ {
		a := neArr[i]
		inciN[h[a]] += xN
		for j := i + 1; j < neN; j++ {
			b := neArr[j]
			inci := h[a] + h[b]
			if c := hpair[a][b]; c > 0 {
				inci -= c
			}
			inciN[inci]++
		}
	}

	// eliminate empty data and sort
	iN := len(inciN)
	inciArr := make([]inciStat, 0, iN)
	for k, v := range inciN {
		inciArr = append(inciArr, inciStat{k, v})
	}
	// reverse sort
	sort.Slice(inciArr, func(i, j int) bool {
		return inciArr[i].inci > inciArr[j].inci
	})

	// get answers
	qN := len(queries)
	res := make([]int, qN)
	qArr := make([]queryInfo, qN)
	for i := 0; i < qN; i++ {
		qArr[i].i = i
		qArr[i].q = queries[i]
	}
	sort.Slice(qArr, func(i, j int) bool {
		return qArr[i].q > qArr[j].q
	})
	prevAcc := 0
	prevJ := 0
	for i := 0; i < qN; i++ {
		for j := prevJ; j < iN; j++ {
			if inciArr[j].inci <= qArr[i].q {
				qArr[i].res = prevAcc
				prevJ = j
				break
			}
			prevAcc += inciArr[j].cnt
		}
	}
	// output res
	for i := 0; i < qN; i++ {
		ri := qArr[i].i
		res[ri] = qArr[i].res
	}
	return res
}

type inciStat struct {
	inci int
	cnt  int
}

type queryInfo struct {
	i   int
	q   int
	res int
}

func map2orderList(neMap map[int]bool) []int {
	res := make([]int, len(neMap))
	for k := range neMap {
		res = append(res, k)
	}
	sort.Ints(res)
	return res
}
