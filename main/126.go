package main

import "math"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	wordMap := make(map[string][]string)
	for _, word := range wordList {
		wordMap[word] = []string{}
	}

	for i := 0; i < len(wordList); i++ {
		for j := i + 1; j < len(wordList); j++ {
			if conn(wordList[i], wordList[j]) {
				wordMap[wordList[i]] = append(wordMap[wordList[i]], wordList[j])
				wordMap[wordList[j]] = append(wordMap[wordList[j]], wordList[i])
			}
		}
	}
	if _, ok := wordMap[beginWord]; !ok {
		for i := 0; i < len(wordList); i++ {
			if conn(beginWord, wordList[i]) {
				wordMap[beginWord] = append(wordMap[beginWord], wordList[i])
			}
		}
	}

	res := make(map[int][][]string)
	minLen := math.MaxInt32
	vis := make(map[string]bool)
	var dfs func(word string, path []string)
	dfs = func(word string, path []string) {
		path = append(path, word)
		if word == endWord {
			l := len(path)
			minLen = min(minLen, l)
			r, ok := res[l]
			if !ok {
				r = [][]string{}
			}
			r = append(r, path)
			res[l] = r
			return
		}

		vis[word] = true
		for _, to := range wordMap[word] {
			if !vis[to] {
				dfs(to, path)
			}
		}
		vis[word] = false
	}
	dfs(beginWord, []string{})
	return res[minLen]
}

func conn(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	if len(a) == 0 {
		return false
	}

	if a[0] == b[0] {
		return conn(a[1:], b[1:])
	}
	return a[1:] == b[1:]
}
