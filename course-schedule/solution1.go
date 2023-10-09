package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	deps := make([]map[int]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		deps[i] = make(map[int]bool)
	}

	dependOn := func(dep, course int) bool {
		cs := map[int]bool{dep: true}
		visitMap := make([]bool, numCourses)
		for len(cs) > 0 {
			nextCs := make(map[int]bool)
			for c := range cs {
				if visitMap[c] {
					continue
				}
				if c == course {
					return true
				}
				visitMap[c] = true
				for d := range deps[c] {
					nextCs[d] = true
				}
			}
			cs = nextCs
		}
		return false
	}

	for i := 0; i < len(prerequisites); i++ {
		c, dep := prerequisites[i][0], prerequisites[i][1]
		if dependOn(dep, c) {
			return false
		}
		deps[c][dep] = true
	}
	return true
}
