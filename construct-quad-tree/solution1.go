package main

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	var helper func(i, j, n int) *Node
	helper = func(i, j, n int) *Node {
		if n == 1 {
			val := true
			if grid[i][j] == 0 {
				val = false
			}
			return &Node{Val: val, IsLeaf: true}
		}
		// try to speed up
		if n == 2 {
			x := grid[i][j]
			if x == grid[i+1][j] &&
				x == grid[i][j+1] &&
				x == grid[i+1][j+1] {
				return &Node{Val: x == 1, IsLeaf: true}
			}
			return &Node{
				TopLeft:     &Node{IsLeaf: true, Val: x == 1},
				TopRight:    &Node{IsLeaf: true, Val: grid[i][j+1] == 1},
				BottomLeft:  &Node{IsLeaf: true, Val: grid[i+1][j] == 1},
				BottomRight: &Node{IsLeaf: true, Val: grid[i+1][j+1] == 1},
			}
		}
		half := n / 2
		tl := helper(i, j, half)
		tr := helper(i, j+half, half)
		bl := helper(i+half, j, half)
		br := helper(i+half, j+half, half)
		if tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf &&
			tl.Val == tr.Val && tl.Val == bl.Val && tl.Val == br.Val {
			return &Node{Val: tl.Val, IsLeaf: true}
		}
		return &Node{
			IsLeaf:      false,
			TopLeft:     tl,
			TopRight:    tr,
			BottomLeft:  bl,
			BottomRight: br,
		}
	}
	res := helper(0, 0, len(grid))
	return res
}
