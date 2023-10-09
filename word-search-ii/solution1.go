package main

/*
func findWords(board [][]byte, words []string) []string {
	root := &Trie{}
	for _, word := range words {
		root.fromWord(word)
	}
	found := make(map[string]bool)

	bHeight := len(board)
	if bHeight <= 0 {
		return nil
	}
	bWidth := len(board[0])
	if bWidth <= 0 {
		return nil
	}
	// var visited map[pos]bool
	visited := make([]bool, bHeight*bWidth)
	// dfs
	var travel func(y, x int, trieN *Trie)
	travel = func(y, x int, trieN *Trie) {
		p := y*bWidth + x
		if visited[p] {
			return
		}
		ch := board[y][x]
		nextTrie := trieN.sons[ch-'a']
		if nextTrie == nil {
			return
		}
		visited[p] = true

		if len(nextTrie.word) > 0 {
			found[string(nextTrie.word)] = true
		}
		if y > 0 {
			travel(y-1, x, nextTrie)
		}
		if y < bHeight-1 {
			travel(y+1, x, nextTrie)
		}
		if x > 0 {
			travel(y, x-1, nextTrie)
		}
		if x < bWidth-1 {
			travel(y, x+1, nextTrie)
		}
		visited[p] = false
	}
	for y := 0; y < bHeight; y++ {
		for x := 0; x < bWidth; x++ {
			travel(y, x, root)
		}
	}

	res := make([]string, 0, len(found))
	for w := range found {
		res = append(res, w)
	}
	return res
}

type Trie struct {
	sons [26]*Trie
	word string
}

func (t *Trie) fromWord(word string) {
	curr := t
	for _, ch := range word {
		c := ch - 'a'
		if curr.sons[c] == nil {
			curr.sons[c] = &Trie{}
		}
		curr = curr.sons[c]
	}
	curr.word = word
}
*/
