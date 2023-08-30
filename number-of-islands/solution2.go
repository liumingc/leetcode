package main

func numIslands(grid [][]byte) int {
    m := len(grid)
    n := len(grid[0])
    mark := make([][]int, m+1)
    for i:=0; i<m+1; i++ {
        mark[i] = make([]int, n+1)
    }
    islN := 0 // ISLand Numbers
    confl := make(map[int]int)
    // confl := make([]int, (m+1) * (n+1))
    for y:=0; y<m; y++ {
        for x:=0; x<n; x++ {
            switch grid[y][x] {
            case '1':
                xPrev := mark[y+1][x]
                yPrev := mark[y][x+1]
                switch {
                case xPrev == 0 && yPrev == 0:
                    // new island
                    islN++
                    mark[y+1][x+1] = islN
                case xPrev == 0:
                    // yPrev != 0
                    mark[y+1][x+1] = yPrev
                case yPrev == 0 || xPrev == yPrev:
                    mark[y+1][x+1] = xPrev
                default:
                    if xPrev > yPrev {
                        xPrev, yPrev = yPrev, xPrev
                    }
                    mark[y+1][x+1] = xPrev
                    for {
                        c := confl[yPrev]
                        confl[yPrev] = xPrev
                        if c == 0 || c == xPrev {
                            break
                        }
                        if c > xPrev {
                            yPrev = c
                        } else {
                            xPrev, yPrev = c, xPrev
                        }
                    }
                }
            }
        }
    }
    printMark(mark)
    islTbl := make(map[int]bool)
    for i:=1; i<=islN; i++ {
        x := i
        for y:=confl[x]; y>0; {
            x = y
            y = confl[x]
        }
        islTbl[x] = true
    }

    return len(islTbl)

}

