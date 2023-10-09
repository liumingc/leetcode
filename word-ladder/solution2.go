package main

func ladderLength2(beginWord string, endWord string, wordList []string) int {
    wordMap := make(map[string]int)
    edges := make(map[int]map[int]bool)
    wID := 0
    addWord := func(x string) int {
        if res, ok := wordMap[x]; ok {
            return res
        }
        wordMap[x] = wID
        edges[wID] = make(map[int]bool)
        wID++
        return wID - 1
    }

    wordN := len(beginWord)
    addEdge := func(x string) {
        xID1 := addWord(x)
        for i:=0; i<wordN; i++ {
            tmp := []byte(x)
            tmp[i] = '*'
            x2 := string(tmp)
            xID2 := addWord(x2)
            edges[xID1][xID2] = true
            edges[xID2][xID1] = true
        }
    }

    for _, word := range wordList {
        addEdge(word)
    }
    if _, ok := wordMap[endWord]; !ok {
        return 0
    }
    addEdge(beginWord)
    // end prepare, now begin to search

    beginID, endID := wordMap[beginWord], wordMap[endWord]
    xs := map[int]bool{beginID: true}
    step := 0
    visited := make([]bool, len(wordMap))
    visited[beginID] = true
    for len(xs) > 0 {
        ys := make(map[int]bool)
        for x := range xs {
            if x == endID {
                return step / 2 + 1
            }

            for e := range edges[x] {
                if visited[e] {
                    continue
                }
                visited[e] = true
                ys[e] = true
            }
        }
        xs = ys
        step++
    }

    return 0
}

