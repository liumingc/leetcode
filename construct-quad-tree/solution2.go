package main

func construct2(grid [][]int) *Node {
	var helper func(i, j, n int)
	helper = func(i, j, n int) {
		if n <= 1 {
			return
		}

		half := n / 2
		helper(i, j, half)
		helper(i, j+half, half)
		helper(i+half, j, half)
		helper(i+half, j+half, half)
		grid[i][j] += grid[i][j+half] +
			grid[i+half][j] +
			grid[i+half][j+half]
	}
	helper(0, 0, len(grid))
	var recurConst func(i, j, n int) *Node
	recurConst = func(i, j, n int) *Node {
		if n == 1 {
			return &Node{
				IsLeaf: true,
				Val:    grid[i][j] > 0,
			}
		}
		if grid[i][j] == 0 {
			return &Node{
				IsLeaf: true,
				Val:    false,
			}
		}
		if grid[i][j] == n*n {
			return &Node{
				IsLeaf: true,
				Val:    true,
			}
		}
		half := n / 2
		grid[i][j] -= grid[i][j+half] +
			grid[i+half][j] +
			grid[i+half][j+half]
		return &Node{
			TopLeft:     recurConst(i, j, half),
			TopRight:    recurConst(i, j+half, half),
			BottomLeft:  recurConst(i+half, j, half),
			BottomRight: recurConst(i+half, j+half, half),
		}
	}
	return recurConst(0, 0, len(grid))
}
