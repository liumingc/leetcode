package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
    wordMap := make(map[string]bool)
    for _, word := range wordList {
        wordMap[word] = true
    }
    if !wordMap[endWord] {
        return 0
    }
    alphabets = "abcdefghijklmnopqrstuvwxyz"

    xs := map[string]bool{beginWord: true}
    step := 1
    for len(xs) > 0 {
        ys := make(map[string]bool)
        for x := range xs {
            if x == endWord {
                return step
            }
            mutateS(x, ys, wordMap)
        }
        step++
        xs = ys
        for y := range ys {
            delete(wordMap, y)
        }
    }

    return 0
}

// if permute changes of word, then there are
// 26 * 10 = 260 choices
// if iterate over wordList, then there are
// 5000 choices at most

var alphabets string

func mutateS(x string, coll map[string]bool, wordMap map[string]bool) map[string]bool {
    for i := 0; i<len(x); i++ {
        tmp := []byte(x)
        for k:=0; k<len(alphabets); k++ {
            if alphabets[k] == tmp[i] {
                continue
            }
            tmp[i] = alphabets[k]
            y := string(tmp)
            if wordMap[y] {
                coll[y] = true
            }
        }
    }
    return coll
}
