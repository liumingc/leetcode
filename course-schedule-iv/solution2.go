package main

func checkIfPrerequisite2(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	isPre := make([][]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		isPre[i] = make([]bool, numCourses)
	}
	inDegree := make([]int, numCourses)
	postMap := make([][]int, numCourses)

	for _, pre := range prerequisites {
		a, b := pre[0], pre[1]
		inDegree[b]++
		postMap[a] = append(postMap[a], b)
	}
	roots := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			roots = append(roots, i)
		}
	}

	for len(roots) > 0 {
		var a int
		a, roots = roots[0], roots[1:]
		for _, b := range postMap[a] {
			isPre[a][b] = true
			inDegree[b]--
			if inDegree[b] == 0 {
				roots = append(roots, b)
			}
			for i := 0; i < numCourses; i++ {
				isPre[i][b] = isPre[i][b] || isPre[i][a]
			}
		}
	}

	qN := len(queries)
	res := make([]bool, qN)
	for i := 0; i < qN; i++ {
		a, b := queries[i][0], queries[i][1]
		res[i] = isPre[a][b]
	}
	return res
}
