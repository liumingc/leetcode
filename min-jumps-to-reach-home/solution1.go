package main

/*
import "fmt"

func minimumJumps(forbidden []int, a int, b int, x int) int {
	forbMap := make([]bool, 2001)
	for i := 0; i < len(forbidden); i++ {
		forbMap[forbidden[i]] = true
	}

	jmps = make([]jmpInfo, 0)

	fN := maxJumps(a, b, x)
	fn := (x + a - 1) / a
	fmt.Printf("maxJumps, fn=%d, fN=%d\n", fn, fN)
	fmt.Println("jmps=", jmps)
	for i := fn; i <= fN; i++ {
		remain := i*a - x
		if remain%b != 0 {
			continue
		}
		bN := remain / b
		ok := tryJump(forbMap, i, bN, a, b, x)
		fmt.Printf("tryJump, fN=%d, bN=%d, x=%d, res=%v\n", i, bN, x, ok)
		if ok {
			return i + bN
		}
	}

	return -1
}

func tryJump(forbMap []bool, fN, bN, a, b, x int) bool {
	var trySolute func(curr int, fN, bN int) bool
	trySolute = func(curr int, fN, bN int) (succ bool) {
		if curr <= 2000 && forbMap[curr] {
			// fmt.Printf("forbid %d\n", curr)
			return false
		}

		if curr == x {
			return true
		}
		if fN == 0 && bN == 0 {
			return false
		}

		// try possible jmp, always forward, never goes back as a whole
		for i := 0; i < len(jmps); i++ {
			ji := jmps[i]
			if fN < ji.fN || bN < ji.bN {
				continue
			}
			// TODO: need to check if intermediate hits forbidden

			if trySolute(curr+ji.c, fN-ji.fN, bN-ji.bN) {
				return true
			}
		}
		return false
	}
	return trySolute(0, fN, bN)
}

func gcd(a, b int) int {
	switch {
	case a > b:
		if a%b == 0 {
			return b
		}
		c := a / b
		a -= b * c
	case a == b:
		return a
	default:
		a, b = b, a
	}
	return gcd(a, b)
}

func maxJumps(a, b, x int) (fN int) {
	if a > b {
		c := a - b
		n := (x + c - 1) / c
		jmps = []jmpInfo{{c, 1, 1}, {a, 1, 0}}
		return n
	}
	if a == b {
		n := (x + a - 1) / a
		jmps = []jmpInfo{{a, 1, 0}}
		return n
	}
	k := gcd(b, a)
	fmt.Printf("a=%d, b=%d, x=%d\n", a, b, x)
	for i := 0; i < a/k; i++ {
		m := i * b / a
		c := a*(m+1) - i*b
		n := (x + c - 1) / c
		jmps = append(jmps, jmpInfo{
			c, m + 1, i,
		})
		fmt.Printf("forw=%d, back=%d, c=%d\n", m+1, i, c)
		tfN := (m + 1) * n
		if tfN > fN {
			fN = tfN
		}
	}
	return fN
}

type jmpInfo struct {
	c      int
	fN, bN int
}

var jmps []jmpInfo
*/
