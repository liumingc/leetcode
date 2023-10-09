package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	res := make([]int, 0)
	deps := make(map[int][]int)
	for _, preReq := range prerequisites {
		course, dep := preReq[0], preReq[1]
		deps[course] = append(deps[course], dep)
	}

	stateMap := make([]courseState, numCourses)
	var dfs func(c int) bool
	dfs = func(c int) bool {
		if stateMap[c] == visited {
			return true
		}
		if stateMap[c] == visiting {
			return false
		}
		stateMap[c] = visiting
		for _, dep := range deps[c] {
			if !dfs(dep) {
				return false
			}
		}
		stateMap[c] = visited
		res = append(res, c)
		return true
	}
	for i := 0; i < numCourses; i++ {
		if stateMap[i] != notVisit {
			continue
		}
		if !dfs(i) {
			return nil
		}
	}
	return res
}

type courseState int8

const (
	notVisit courseState = iota
	visiting
	visited
)
