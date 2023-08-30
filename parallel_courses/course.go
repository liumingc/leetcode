package main

func minimumTime(n int, relations [][]int, time []int) int {
	// build graph
	tbl := make(map[int]*Course)
	for i := 0; i < n; i++ {
		c := &Course{
			nr:       i,
			selfTime: time[i],
		}
		tbl[i] = c
	}

	for _, rel := range relations {
		previ, nexti := rel[0]-1, rel[1]-1
		prev, next := tbl[previ], tbl[nexti]
		next.prev = append(next.prev, prev)
		prev.next = next
	}
	// find root
	var roots []*Course
	for _, c := range tbl {
		if c.next == nil {
			roots = append(roots, c)
		}
	}
	minTime := 0
	for _, r := range roots {
		t := r.needTime()
		if t > minTime {
			minTime = t
		}
	}
	return minTime
}

type Course struct {
	nr       int
	selfTime int
	// visit bool
	accTime int
	prev    []*Course
	next    *Course
}

func (c *Course) needTime() int {
	if c.accTime > 0 {
		return c.accTime
	}
	prevTime := 0
	for _, prev := range c.prev {
		nt := prev.needTime()
		if nt > prevTime {
			prevTime = nt
		}
	}
	c.accTime = prevTime + c.selfTime
	return c.accTime
}

/*
func minimumTime(n int, relations [][]int, time []int) int {
    relTbl := make(map[int][]int, len(relations))
    for _, rel := range relations {
        prev, next := rel[0]-1, rel[1]-1
        relTbl[next] = append(relTbl[next], prev)
    }

    accTbl = make([]int, n)
    minTime := 0
    for i:=0; i<n; i++ {
        x := getTime(i, relTbl, time)
        if x > minTime {
            minTime = x
        }
    }
    return minTime
}

var accTbl []int
func getTime(c int, relTbl map[int][]int, time []int) int {
    if accTbl[c] > 0 {
        return accTbl[c]
    }
    depTime := 0
    for _, prev := range relTbl[c] {
        ct := getTime(prev, relTbl, time)
        if ct > depTime {
            depTime = ct
        }
    }
    accTbl[c] = depTime + time[c]
    return accTbl[c]
}
*/

// faster
/*
func minimumTime(n int, relations [][]int, time []int) int {
    relTbl := make([][]int, n)
    for _, rel := range relations {
        prev, next := rel[0]-1, rel[1]-1
        relTbl[next] = append(relTbl[next], prev)
    }

    accTbl = make([]int, n)
    minTime := 0
    for i:=0; i<n; i++ {
        x := getTime(i, relTbl, time)
        if x > minTime {
            minTime = x
        }
    }
    return minTime
}

var accTbl []int
func getTime(c int, relTbl [][]int, time []int) int {
    if accTbl[c] > 0 {
        return accTbl[c]
    }
    depTime := 0
    for _, prev := range relTbl[c] {
        ct := getTime(prev, relTbl, time)
        if ct > depTime {
            depTime = ct
        }
    }
    accTbl[c] = depTime + time[c]
    return accTbl[c]
}
*/
