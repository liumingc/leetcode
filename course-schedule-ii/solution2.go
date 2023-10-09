package main

func findOrder2(numCourses int, prerequisites [][]int) []int {
	res := make([]int, 0)
	depNumMap := make([]int16, numCourses)
	roots := make(map[int]bool)
	nonRoots := make(map[int]bool)
	postMap := make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		roots[i] = true
	}
	for _, prereq := range prerequisites {
		c, dep := prereq[0], prereq[1]
		postMap[dep] = append(postMap[dep], c)
		depNumMap[c]++
		if depNumMap[c] == 1 {
			delete(roots, c)
			nonRoots[c] = true
		}
	}

	for len(roots) > 0 {
		nxtRoots := make(map[int]bool)
		for root := range roots {
			res = append(res, root)
			for _, post := range postMap[root] {
				depNumMap[post]--
				if depNumMap[post] == 0 {
					nxtRoots[post] = true
					delete(nonRoots, post)
				}
			}
		}
		roots = nxtRoots
	}
	if len(nonRoots) == 0 {
		return res
	}

	return nil
}
