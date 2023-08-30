package main

func findSubstring2(s string, words []string) []int {
	wl := len(words)
	if wl < 1 {
		return nil
	}

	wil := len(words[0])
	joinL := wl * wil
	res := make([]int, 0)
	tbl := buildFastTable2(words)

	for i := 0; i <= len(s)-joinL; i++ {
		seen := make(map[int]bool)

		for w := 0; w < wl; w++ {
			wstart := i + w*wil
			choices := tbl[s[wstart]]
			if len(choices) == 0 {
				break
			}
			match := false
			for _, choice := range choices {
				if seen[choice] {
					continue
				}
				word := words[choice]
				// compare word
				match = true
				for p := 1; p < wil; p++ {
					if word[p] != s[wstart+p] {
						match = false
						break
					}
				}
				if match {
					// fmt.Printf("i=%d, w=%d,choices=%d\n", i, w, choice)
					seen[choice] = true
					break
				}
			}
			if !match {
				break
			} else {
				if w == wl-1 {
					res = append(res, i)
				}
			}
		}
	}

	return res
}

func buildFastTable2(words []string) [][]int {
	res := make([][]int, 256)
	for i, word := range words {
		c := word[0]
		res[c] = append(res[c], i)
	}
	return res
}
