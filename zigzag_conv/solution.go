package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	res := zoutput(s, numRows)
	show(res, numRows)
	return reconstruct(res, numRows)
}

func zoutput(s string, numRows int) [][]byte {
	zy := 0
	// zx := 0
	res := [][]byte{}
	var col []byte
	goZ := false

	for i := 0; i < len(s); i++ {
		if zy == 0 || goZ {
			col = make([]byte, numRows)
			res = append(res, col)
		}

		// fmt.Printf("char [%d] = %c => zy=[%d]\n", i, s[i], zy)
		col[zy] = s[i]

		// move to next position
		if goZ {
			zy--
			// zx++
			if zy == 0 {
				goZ = false
			} else if zy == -1 {
				goZ = false
				zy = 1
			}
		} else {
			if zy == numRows-1 {
				if zy == 0 {
					// zx++ // special numRows
					continue
				}
				goZ = true
				zy--
				// zx++
			} else {
				zy++
			}
		}

	}
	return res
}

func show(arr [][]byte, numRows int) {
	xEnd := len(arr)
	yEnd := numRows
	fmt.Printf("%d X %d\n", xEnd, yEnd)
	/*
		for _, col := range arr {
			fmt.Printf("%s\n", col)
		}
		fmt.Println("==>")
	*/
	for x := 0; x < xEnd; x++ {
		for y := 0; y < yEnd; y++ {
			c := arr[x][y]
			if c == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("%c", c)
			}
		}
		fmt.Printf("\n")
	}
}

func reconstruct(arr [][]byte, numRows int) string {
	var sb strings.Builder
	xEnd := len(arr)
	yEnd := numRows
	for x := 0; x < xEnd; x++ {
		for y := 0; y < yEnd; y++ {
			c := arr[x][y]
			if c == 0 {
				continue
			}
			sb.WriteByte(c)
		}
	}
	return sb.String()
}

/*
func convert(s string, numRows int) string {
    res := zoutput(s, numRows)
    return reconstruct(res, numRows)
}

func zoutput(s string, numRows int) [][]byte {
    zy := 0
    // zx := 0
    res := [][]byte{}
    var col []byte
    goZ := false

    for i:=0; i<len(s); i++ {
        if zy == 0 || goZ {
            col = make([]byte, numRows)
            res = append(res, col)
        }

        col[zy] = s[i]

        // move to next position
        if goZ {
            if zy == 0 {
                goZ = false
                zy++
            } else {
                zy--
            }
        } else {
            if zy == numRows-1 {
                if zy == 0 {
                    // zx++
                    continue
                }
                goZ = true
                zy--
                // zx++
            } else {
                zy++
            }
        }

    }
    return res
}

func reconstruct(arr [][]byte, numRows int) string {
    var sb strings.Builder
    xEnd := len(arr)
    yEnd := numRows
    for y:=0; y<yEnd;y++ {
        for x:=0; x<xEnd; x++ {
            c := arr[x][y]
            if c == 0 {
                continue
            }
            sb.WriteByte(c)
        }
    }
    return sb.String()
}
*/
