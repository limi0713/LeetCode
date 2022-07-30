package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

// BenchmarkCopy copy
func BenchmarkCopy(b *testing.B) {
	numsO := make([]int, 10000)
	// numsT := make([]int, len(numsO), len(numsO))
	for i := 0; i < b.N; i++ {
		numsT := make([]int, len(numsO), len(numsO))
		copy(numsT, numsO)
	}
}

// BenchmarkAppend append
func BenchmarkAppend(b *testing.B) {
	numsO := make([]int, 10000)
	numsT := make([]int, 0, len(numsO))
	for i := 0; i < b.N; i++ {
		numsT = append(numsT[:0:0], numsO...)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge1(p1, p2 *ListNode) *ListNode {
	if p1 == nil {
		return p2
	}
	if p2 == nil {
		return p1
	}

	p := &ListNode{}
	head := p
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = &ListNode{
				Val: p2.Val,
			}
			p = p.Next
			p2 = p2.Next
		} else {
			p.Next = &ListNode{
				Val: p1.Val,
			}
			p = p.Next
			p1 = p1.Next
		}
	}

	for p1 != nil {
		p.Next = &ListNode{
			Val: p1.Val,
		}
		p = p.Next
		p1 = p1.Next
	}

	for p2 != nil {
		p.Next = &ListNode{
			Val: p2.Val,
		}
		p = p.Next
		p2 = p2.Next
	}
	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for {
		mRet := make([]*ListNode, 0)
		if len(lists) == 1 {
			return lists[0]
		}

		j, _ := json.Marshal(lists)
		fmt.Printf("list: %+v\n", string(j))

		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) {
				mRet = append(mRet, merge1(lists[i], lists[i+1]))
			} else {
				mRet = append(mRet, lists[i])
			}
		}
		lists = append(make([]*ListNode, 0), mRet...)
	}
}

func TestMerge(t *testing.T) {
	lists := []*ListNode{
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 6,
				},
			},
		},
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			}},
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
			}},
	}

	r := mergeKLists(lists)

	j, _ := json.Marshal(r)
	fmt.Printf("%+v\n", string(j))
	time.Sleep(time.Second)
}
