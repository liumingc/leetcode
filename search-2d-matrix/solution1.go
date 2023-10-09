package main

func searchMatrix(matrix [][]int, target int) bool {
	rN := len(matrix)
	cN := len(matrix[0])

	bsearch := func(row []int, target int) bool {
		low := 0
		high := cN - 1
		if target < row[low] {
			return false
		}
		if target == row[low] {
			return true
		}
		for low < high {
			mid := (low + high) / 2
			if target == row[mid] {
				return true
			}
			if target > row[mid] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		return row[low] == target
	}

	for y := 0; y < rN; y++ {
		if matrix[y][cN-1] < target {
			continue
		}
		return bsearch(matrix[y], target)
	}
	return false
}
