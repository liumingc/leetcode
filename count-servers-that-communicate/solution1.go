package main

func countServers(grid [][]int) int {
	visited := make(map[pos]bool)

	// check rows
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	for y := 0; y < m; y++ {
		cells := make([]int, 0, n)
		for x := 0; x < n; x++ {
			if grid[y][x] == 1 {
				cells = append(cells, x)
			}
		}
		if len(cells) > 1 {
			for i := 0; i < len(cells); i++ {
				visited[pos{y, cells[i]}] = true
			}
		}
	}
	// check columns
	for x := 0; x < n; x++ {
		cells := make([]int, 0, m)
		for y := 0; y < m; y++ {
			if grid[y][x] == 1 {
				cells = append(cells, y)
			}
		}
		if len(cells) > 1 {
			for i := 0; i < len(cells); i++ {
				visited[pos{cells[i], x}] = true
			}
		}
	}
	return len(visited)
}

type pos struct {
	x, y int
}
