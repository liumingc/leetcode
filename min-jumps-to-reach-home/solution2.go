package main

func minimumJumps(forbidden []int, a int, b int, x int) int {
	maxForb := forbidden[0]
	forbMap := make(map[int]bool, len(forbidden))
	for _, fb := range forbidden {
		if fb > maxForb {
			maxForb = fb
		}
		forbMap[fb] = true
	}

	lower := 0
	upper := max(maxForb+a, x) + b
	arr := []*elem{
		{0, true, 0},
	}
	visitMap := map[velem]bool{
		{0, true}: true,
	}

	for len(arr) > 0 {
		nextArr := make([]*elem, 0)
		for _, e := range arr {
			if e.pos == x {
				return e.step
			}
			// try right
			nextPos := e.pos + a
			ve := velem{nextPos, true}
			if nextPos <= upper &&
				!forbMap[nextPos] &&
				!visitMap[ve] {
				visitMap[ve] = true
				nextArr = append(nextArr, &elem{nextPos, true, e.step + 1})
			}
			// try left
			if e.right { // only if previous is right
				nextPos = e.pos - b
				ve = velem{nextPos, false}
				if nextPos >= lower &&
					!forbMap[nextPos] &&
					!visitMap[ve] {
					visitMap[ve] = true
					nextArr = append(nextArr, &elem{nextPos, false, e.step + 1})
				}
			}
		}
		arr = nextArr
	}
	return -1
}

type elem struct {
	pos   int
	right bool // prev direction
	step  int
}

type velem struct {
	pos   int
	right bool
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
