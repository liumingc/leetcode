package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)
	n12 := n1 + n2
	needTwo := false
	pos := n12 / 2
	if n12&1 == 0 {
		needTwo = true
		pos--
	}
	median := func(arr []int, pos int) float64 {
		if needTwo {
			return float64(arr[pos]+arr[pos+1]) / 2
		}
		return float64(arr[pos])
	}

	switch {
	case n1 == 0:
		return median(nums2, pos)
	case n2 == 0:
		return median(nums1, pos)
	}

	bsearch := func(arr []int, l, h, x int) int {
		/*
			if l >= len(arr) {
				return l
			}
			if h >= len(arr) {
				return h
			}
		*/
		for l <= h {
			m := (l + h) / 2
			if arr[m] == x {
				return m
			}
			if arr[m] < x {
				l = m + 1
			} else {
				h = m - 1
			}
		}
		return l
	}

	getNext := func(arr1, arr2 []int, i1, i2 int) int {
		if i1 >= len(arr1) {
			return arr2[i2]
		}
		if i2 >= len(arr2) {
			return arr1[i1]
		}
		return min(arr1[i1], arr2[i2])
	}

	l1, h1 := 0, n1-1
	l2, h2 := 0, n2-1
	m1, m2 := 0, 0
	for l1 <= h1 {
		m1 = (l1 + h1) / 2
		fmt.Printf("bsearch arr2, l2=%d, h2=%d\n", l2, h2)
		m2 = bsearch(nums2, l2, h2, nums1[m1])
		fmt.Printf(".. m2=%d\n", m2)
		m12 := m1 + m2
		if m12 == pos {
			fmt.Printf("find in m1, m1=%d, m2=%d\n", m1, m2)
			if needTwo {
				nxt := getNext(nums1, nums2, m1+1, m2)
				return float64(nums1[m1]+nxt) / 2
			}
			return float64(nums1[m1])
		}
		if m12 < pos {
			l1 = m1 + 1
			l2 = m2
		} else {
			h1 = m1 - 1
			if m2 < n2 {
				h2 = m2
			}
		}
	}
	pos = pos - l1
	fmt.Printf("find after m1, l1=%d, h1=%d, pos=%d\n", l1, h1, pos)
	if !needTwo {
		return float64(nums2[pos])
	}
	nxt := getNext(nums2, nums1, pos+1, l1)
	return float64(nums2[pos]+nxt) / 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
