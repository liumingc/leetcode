package main

/*
func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	wl := len(words)
	if wl < 1 {
		return nil
	}

	wil := len(words[0])
	tot := len(s)
	c2w, w2n := buildFastTable(words)

	joinL := wl * wil
	lower := tot - joinL
	if lower < 0 {
		return nil
	}

	// 0 .. wil-1
	pos2word := make([]int, lower+1)
	word2pos := make([][][]int, wil)
	for i := 0; i < wil; i++ {
		word2pos[i] = make([][]int, wl)
	}
	for i := 0; i <= lower; i++ {
		pos2word[i] = -1
	}

	handleOne := func(offset, i, matchN int, late bool) int {
		wstart := i

		if prev := i - joinL; late && prev >= 0 {
			wi := pos2word[prev]
			if prev >= 0 && wi >= 0 {
				if word2pos[offset][wi][0] == prev {
					word2pos[offset][wi] = word2pos[offset][wi][1:]
					matchN--
				}
			}
		}

		choices := c2w[s[wstart]]
		if len(choices) == 0 {
			return matchN
		}

		for _, wi := range choices {
			word := words[wi]
			found := true
			for p := 1; p < wil; p++ {
				if word[p] != s[wstart+p] {
					found = false
					break
				}
			}
			if found {
				if wstart <= lower {
					pos2word[wstart] = wi
				}
				xs := word2pos[offset][wi]
				maxN := w2n[wi]
				if len(xs) == maxN {
					// retire the first one
					word2pos[offset][wi] = append(xs[1:], wstart)
				} else {
					word2pos[offset][wi] = append(xs, wstart)
					matchN++
				}
				return matchN
			}
		}
		return matchN
	}

	matchTbl := make([]int, wil)
	for i := 0; i < wil && i <= len(s)-joinL; i++ {
		for k := 0; k < wl; k++ {
			matchTbl[i] = handleOne(i, i+k*wil, matchTbl[i], false)
		}
		if matchTbl[i] == wl {
			res = append(res, i)
		}
	}

	for i, offset := wil, 0; i <= len(s)-joinL; i++ {
		if offset == wil {
			offset = 0
		}
		matchTbl[offset] = handleOne(offset, i+joinL-wil, matchTbl[offset], true)
		if matchTbl[offset] == wl {
			res = append(res, i)
		}
		offset++
	}

	return res
}

func buildFastTable(words []string) (c2w [][]int, w2n []int) {
	c2w = make([][]int, 256)
	w2n = make([]int, len(words))
	m := make(map[string]int)
	for i, word := range words {
		x, ok := m[word]
		if !ok {
			m[word] = i
			w2n[i] = 1
			continue
		}
		w2n[x] = w2n[x] + 1
	}
	for wi := range w2n {
		c := words[wi][0]
		c2w[c] = append(c2w[c], wi)
	}
	return
}
*/
