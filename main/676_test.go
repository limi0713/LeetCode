package main

import (
	"fmt"
	"testing"
)

func Test676(t *testing.T) {
	b := Constructor676()
	b.BuildDict([]string{"hello", "leetcode"})
	res := b.Search("hello")
	fmt.Print(res)
}

func Test3137(t *testing.T) {
	minimumOperationsToMakeKPeriodic("leetcoleet", 2)
}

func Test2908(t *testing.T) {
	minimumSum([]int{8, 6, 1, 5, 3})
}

func TestXYJ(t *testing.T) {
	x := 0
	n := 3
	j := 1
	i := 1
	x |= n >> j & (1 << i)
	fmt.Printf("x = %d", x)
}

func Test3133(t *testing.T) {
	minEnd(3, 4)
}

func Test126(t *testing.T) {
	beginWord := "red"
	endWord := "tax"
	wordList := []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}

	_ = findLadders(beginWord, endWord, wordList)
}

func Test148(t *testing.T) {
	head := initList([]int{4, 2, 1, 3})
	_ = sortList(head)
}

func initList(nums []int) *ListNode {
	start := &ListNode{}
	p := start
	for _, num := range nums {
		p.Next = &ListNode{
			Val:  num,
			Next: nil,
		}
		p = p.Next
	}
	return start.Next
}

func Test211(t *testing.T) {
	d := Constructor211()
	d.AddWord("bad")
	d.AddWord("dad")
	d.AddWord("mad")
	d.Search("pad")
	d.Search("bad")
	d.Search(".ad")
	d.Search("b..")
}
