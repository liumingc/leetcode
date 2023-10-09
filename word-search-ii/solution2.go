package main

func findWords(board [][]byte, words []string) []string {
	root := &Trie{}
	for _, word := range words {
		root.fromWord(word)
	}

	bHeight := len(board)
	if bHeight <= 0 {
		return nil
	}
	bWidth := len(board[0])
	if bWidth <= 0 {
		return nil
	}
	// var visited map[pos]bool
	res := make([]string, 0)
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
			res = append(res, nextTrie.word)
			nextTrie.word = ""
			if nextTrie.sonCount == 0 {
				trieN.sons[ch-'a'] = nil
				trieN.sonCount--
			}
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

	return res
}

type Trie struct {
	sons     [26]*Trie
	word     string
	sonCount int
}

func (t *Trie) fromWord(word string) {
	curr := t
	for _, ch := range word {
		c := ch - 'a'
		if curr.sons[c] == nil {
			curr.sons[c] = &Trie{}
			curr.sonCount++
		}
		curr = curr.sons[c]
	}
	curr.word = word
}
