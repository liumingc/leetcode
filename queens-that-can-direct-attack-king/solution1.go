package main

func queensAttacktheKing(queens [][]int, king []int) [][]int {
    board := [8][8]bool{}
    //qN := len(queens)
    for _, q := range queens {
        x, y := q[0], q[1]
        board[y][x] = true
    }
    res := make([][]int, 0)
    kx, ky := king[0], king[1]
    // row
    for i:=kx-1; i>=0; i-- {
        if board[ky][i] {
            res = append(res, []int{i, ky})
            break
        }
    }
    for i:=kx+1; i<8; i++ {
        if board[ky][i] {
            res = append(res, []int{i, ky})
            break
        }
    }
    // col
    for i:=ky-1; i>=0; i-- {
        if board[i][kx] {
            res = append(res, []int{kx, i})
            break
        }
    }
    for i:=ky+1; i<8; i++ {
        if board[i][kx] {
            res = append(res, []int{kx, i})
            break
        }
    }
    // left diag
    for y, x := ky-1, kx-1; x >= 0 && y >= 0; {
        if board[y][x] {
            res = append(res, []int{x, y})
            break
        }
        x--
        y--
    }
    for y, x := ky+1, kx+1; x <8 && y < 8; {
        if board[y][x] {
            res = append(res, []int{x, y})
            break
        }
        x++
        y++
    }
    // right diag
    for y, x := ky-1, kx+1; y >= 0 && x < 8; {
        if board[y][x] {
            res = append(res, []int{x, y})
            break
        }
        y--
        x++
    }
    for y, x := ky+1, kx-1; y <8 && x >= 0; {
        if board[y][x] {
            res = append(res, []int{x, y})
            break
        }
        x--
        y++
    }
    return res
}
