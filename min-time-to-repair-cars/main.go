package main

import "fmt"

func main() {
	ranks := []int{
		5, 1, 8,
	}
	cars := 6
	res := repairCars(ranks, cars)
	fmt.Println(res)
	for i := 15; i <= 18; i++ {
		fmt.Println(i, repairN_inMin(ranks, int64(i)))
	}
}
