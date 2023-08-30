package main

/*
func numIslands(grid [][]byte) int {
	// an island is represented as a set of nodes
	isls := make([]map[node]bool, 0)
	m := len(grid)
	n := len(grid[0])
	konnIsl := func(nx node) (res map[node]bool) {
		islN := len(isls)
		join := false
		for i := islN - 1; i >= 0; i-- {
			isl := isls[i]
			for n := range isl {
				if isKonn(n, nx) {
					if res == nil {
						res = isl
					} else {
						join = true
						// need to join islands
						for n := range isl {
							res[n] = true
						}
						isls[i] = nil
					}
					break
				}
			}
		}
		if join {
			p, q := 0, islN-1
			for ; p < q; p++ {
				if isls[p] == nil {
					isls[p] = isls[q]
					q--
				}
			}
			isls = isls[0 : q+1]
		}
		return res
	}

	var currIsl map[node]bool
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			switch grid[y][x] {
			case '1':
				// has connection with previous islands?
				nx := node{x, y}
				if isl := konnIsl(nx); isl != nil {
					isl[nx] = true
				} else {
					currIsl = map[node]bool{nx: true}
					isls = append(isls, currIsl)
				}
			}
		}
	}
	return len(isls)
}

type node struct {
	x, y int
}

// n2 is the later node
func isKonn(n1, n2 node) bool {
	if n1.x == n2.x && n2.y-n1.y == 1 {
		return true
	}
	if n1.y == n2.y && n2.x-n1.x == 1 {
		return true
	}
	return false
}
*/
