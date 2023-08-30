package main

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	// fast test
	tot := 0
	for i := 0; i < n; i++ {
		tot += gas[i] + cost[i]
	}
	if tot < 0 {
		return -1
	}

	tryGo := func(p int) int {
		curr := gas[p]
		org := p
		for {
			curr -= cost[p]
			if curr < 0 {
				return -1
			}
			p++
			if p == n {
				p = 0
			}
			if p == org {
				return org
			}
			curr += gas[p]
		}
	}
	prevInc := false
	for i := 0; i < n; i++ {
		if prevInc {
			prevInc = gas[i]-cost[i] >= 0
			continue
		}
		res := tryGo(i)
		if res >= 0 {
			return res
		}
		prevInc = gas[i]-cost[i] >= 0
	}
	return -1
}
