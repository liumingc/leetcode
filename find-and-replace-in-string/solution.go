package main

import (
	"sort"
	"strings"
)

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	k := len(indices)
	n := len(s)

	subs := make([]*repl, 0, k)
	sorted := true
	for i := 0; i < k-1; i++ {
		if sorted && indices[i] > indices[i+1] {
			sorted = false
		}
		subs = append(subs, &repl{indices[i], sources[i], targets[i]})
	}
	last := k - 1
	subs = append(subs, &repl{indices[last], sources[last], targets[last]})
	if !sorted {
		sort.Slice(subs, func(i, j int) bool {
			return subs[i].i < subs[j].i
		})
	}

	start := 0
	var sb strings.Builder
	for i := 0; i < k; i++ {
		bN := subs[i].i
		if start < bN {
			sb.WriteString(s[start:bN]) // copy previous
			start = bN
		}
		kn := len(subs[i].src)
		// check if match
		match := false
		eN := bN + kn
		if eN <= n {
			if s[bN:eN] == subs[i].src {
				match = true
			}
		} else {
			eN = n
		}
		if match {
			sb.WriteString(subs[i].tgt)
			start = eN
		}
	}
	if start < n {
		sb.WriteString(s[start:n])
	}

	return sb.String()
}

type repl struct {
	i   int
	src string
	tgt string
}
