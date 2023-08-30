package main

/*
func convert(s string, numRows int) string {
    var sb strings.Builder
    rows := make([][]byte, numRows)
    y := 0
    z := false
    for i:=0; i<len(s); i++ {
        rows[y] = append(rows[y], s[i])

        // move position
        if z == true {
            if y == 0 {
                y = 1
                z = false
            } else {
                y--
            }
        } else {
            if y == numRows - 1 {
                if y > 0 {
                    y--
                    z = true
                }
            } else {
                y++
            }
        }
    }
    for i:=0; i<numRows; i++ {
        sb.Write(rows[i])
    }
    return sb.String()
}

*/
