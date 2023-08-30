package main

import "fmt"

func ways(pizza []string, k int) int {
	piz = pizza
	rN = len(pizza)
	cN = len(piz[0])
	cache = make(map[keyT]*valueT)
	applesN := applesMatN(0, 0)
	fmt.Println("apples = ", applesN)
	res := kutN(0, 0, applesN, k)
	return int(res % (1e9 + 7))
}

var piz []string
var rN int
var cN int

// return apple numbers in row X, starting from col Y
func applesRowN(x0, y0 int) int {
	res := 0
	for x := x0; x < cN; x++ {
		if piz[y0][x] == 'A' {
			res++
		}
	}
	return res
}

func applesColN(x0, y0 int) int {
	res := 0
	for y := y0; y < rN; y++ {
		if piz[y][x0] == 'A' {
			res++
		}
	}
	return res
}

func applesMatN(x0, y0 int) int {
	res := 0
	for y := y0; y < rN; y++ {
		res += applesRowN(x0, y)
	}
	return res
}

type keyT struct {
	x, y int
}

type resT int64

type valueT struct {
	minK int          // impossible K
	tbl  map[int]resT // k |-> number of ways
}

var cache map[keyT]*valueT

func kutN(x0, y0, applesN, k int) (res resT) {
	// check cache first
	key := keyT{x0, y0}
	val, ok := cache[key]
	if ok {
		if k >= val.minK {
			return 0
		}
		res, ok = val.tbl[k]
		if ok {
			return res
		}
	} else {
		val = &valueT{
			minK: applesN + 1,
			tbl:  make(map[int]resT),
		}
		cache[key] = val
	}

	// update cache
	defer func() {
		val.tbl[k] = res
		if res == 0 && k < val.minK {
			val.minK = k
		}
	}()

	switch {
	case applesN < k:
		return
	case k == 1:
		return 1
	/*
		case x0 == rN - 1:
			// last row
			res = kutOne(applesN, k)
		case y0 == cN - 1:
			// last col
			res = kutOne(applesN, k)
	*/
	default:
		// try rows
		dec := 0
		for y := y0; y < rN-1; y++ {
			apples := applesRowN(x0, y)
			if apples == 0 && dec == 0 {
				continue
			}
			dec += apples
			t := kutN(x0, y+1, applesN-dec, k-1)
			if t == 0 {
				break
			}
			res += t
		}
		// try cols
		dec = 0
		for x := x0; x < cN-1; x++ {
			apples := applesColN(x, y0)
			if apples == 0 && dec == 0 {
				continue
			}
			dec += apples
			t := kutN(x+1, y0, applesN-dec, k-1)
			if t == 0 {
				break
			}
			res += t
		}
	}
	return
}

/*
type nk struct {
	n, k int
}
var ctbl map[nk]int

func kutOne(n, k int) int {
	if n < k {
		return 0
	}
	res, ok := ctbl[nk{n, k}]
	if ok {
		return res
	}
	kexcl := 1
	nexcl := 1
	x := n - 1
	for i:=1; i<=k-1; i++ {
		kexcl *= i
		nexcl *= x
		x--
	}
	res = nexcl / kexcl
	ctbl[nk{n, k}] = res
	return res
}
*/
