package main

func splitNum2(num int) int {
	arr := [10]int{}
	for num > 0 {
		low := num % 10
		num = num / 10
		arr[low]++
	}
	num1, num2 := 0, 0
	selFirst := true
	for i := 0; i < 10; {
		if arr[i] == 0 {
			i++
			continue
		}
		arr[i]--
		if selFirst {
			num1 = num1*10 + i
			selFirst = false
		} else {
			num2 = num2*10 + i
			selFirst = true
		}
	}

	return num1 + num2
}
