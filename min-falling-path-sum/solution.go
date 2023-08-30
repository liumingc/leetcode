func minFallingPathSum(grid [][]int) int {
    n := len(grid)
    if n == 1 {
        return grid[0][0]
    }
    tbl = make(map[int]*triple, n)
    cache = make(map[rk]int)
    tri := findMins(grid, 0)
    return tryFall(grid, 0, tri, -1)
}

type rk struct {
    row int
    prev int
}
var cache map[rk]int

func tryFall(grid [][]int, rowN int, tri *triple, prev int) (res int) {
    if rowN == len(grid) - 1 {
        // no next row
        if tri.x1 != prev {
            return grid[rowN][tri.x1]
        }
        return grid[rowN][tri.x2]
    }
    rk := rk{rowN, prev}
    var ok bool
    res, ok = cache[rk]
    if ok {
        return res
    }
    defer func() {
        if len(grid) - rowN > 4 {
            cache[rk] = res
        }
    }()

    nextRow := rowN + 1
    nextTri := findMins(grid, nextRow)
    if tri.x1 == prev {
        // only can choose x2 or x3
        sum2 := tryFall(grid, nextRow, nextTri, tri.x2) + grid[rowN][tri.x2]
        if !conflictWith(tri.x2, nextTri) {
            res = sum2
            return
        }
        sum3 := tryFall(grid, nextRow, nextTri, tri.x3) + grid[rowN][tri.x3]
        res = min(sum2, sum3)
        return res
    }
    sum1 := tryFall(grid, nextRow, nextTri, tri.x1) + grid[rowN][tri.x1]
    if !conflictWith(tri.x1, nextTri) {
        res = sum1
        return
    }
    sum2 := tryFall(grid, nextRow, nextTri, tri.x2) + grid[rowN][tri.x2]
    if !conflictWith(tri.x2, nextTri) {
        res = min(sum1, sum2)
        return res
    }
    sum3 := tryFall(grid, nextRow, nextTri, tri.x3) + grid[rowN][tri.x3]
    res = min(sum1, min(sum2, sum3))

    return res
}

func conflictWith(i int, tri *triple) bool {
    switch {
    case i == tri.x1:
        return true
    case i == tri.x2:
        return true
    case tri.x3 != -1 && i == tri.x3:
        return true
    }
    return false
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

type triple struct {
    x1, x2, x3 int
}

var tbl map[int]*triple
func findMins(grid [][]int, rowN int) *triple {
    res := tbl[rowN]
    if res != nil {
        return res
    }
    x1, x2, x3 := 0, 1, -1
    if grid[rowN][1] < grid[rowN][0] {
        x1, x2 = 1, 0
    }
    for i:=2; i<len(grid); i++ {
        if grid[rowN][i] < grid[rowN][x1] {
            x3, x2, x1 = x2, x1, i
        } else if grid[rowN][i] < grid[rowN][x2] {
            x2, x3 = i, x2
        } else if x3 == -1 || grid[rowN][i] < grid[rowN][x3] {
            x3 = i
        }
    }
    res = &triple{
        x1, x2, x3,
    }
    tbl[rowN] = res
    return res
}
