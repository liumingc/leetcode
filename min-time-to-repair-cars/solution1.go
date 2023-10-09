package main

import (
	"fmt"
	"math"
)

func repairCars(ranks []int, cars int) int64 {
	if cars <= 0 {
		return 0
	}
	if len(ranks) == 0 {
		return math.MaxInt
	}
	rMin := ranks[0]
	// rMax := ranks[0]
	for i := 1; i < len(ranks); i++ {
		if ranks[i] < rMin {
			rMin = ranks[i]
		}
	}

	repairN := func(t int64) int {
		res := 0
		for i := 0; i < len(ranks); i++ {
			res += int(math.Floor(math.Sqrt(float64(t) / float64(ranks[i]))))
			if res > cars {
				return res
			}
		}
		return res
	}
	// sort.Ints(ranks)
	high := int64(cars) * int64(cars) * int64(rMin)
	low := int64(1)
	for low < high {
		mid := (low + high) / 2
		n := repairN(mid)
		fmt.Printf("[%d, %d], %d => %d\n", low, high, mid, n)
		if n >= cars {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func repairN_inMin(ranks []int, t int64) int {
	res := 0
	for i := 0; i < len(ranks); i++ {
		res += int(math.Floor(math.Sqrt(float64(t) / float64(ranks[i]))))
	}
	return res
}
