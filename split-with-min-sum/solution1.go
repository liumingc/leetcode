package main

// https://leetcode.cn/problems/split-with-minimum-sum/solutions/2470641/zui-xiao-he-fen-ge-by-leetcode-solution-6fde/

func splitNum(num int) int {
	arr := []int{}
	for num > 0 {
		low := num % 10
		rem := num / 10
		arr = append(arr, low)
		num = rem
	}
	sort.Ints(arr)
	num1, num2 := 0, 0
	for len(arr) >= 2 {
		a, b := arr[0], arr[1]
		arr = arr[2:]
		num1 = num1*10 + a
		num2 = num2*10 + b
	}
	if len(arr) > 0 {
		num1 = num1*10 + arr[0]
	}
	return num1 + num2
}
