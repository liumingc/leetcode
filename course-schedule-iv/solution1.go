package main

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	preMap := make(map[int]map[int]bool)
	postMap := make(map[int]map[int]bool)
	for i := 0; i < numCourses; i++ {
		preMap[i] = make(map[int]bool)
		postMap[i] = make(map[int]bool)
	}
	for _, pre := range prerequisites {
		a, b := pre[0], pre[1]
		for post := range preMap[b] {
			preMap[a][post] = true
			postMap[post][a] = true
		}
		for f := range postMap[a] {
			preMap[f][b] = true
			postMap[b][f] = true
		}
		preMap[a][b] = true
		postMap[b][a] = true
	}

	res := make([]bool, len(queries))
	for i, query := range queries {
		a, b := query[0], query[1]
		res[i] = preMap[a][b]
	}
	return res
}
