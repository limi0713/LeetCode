package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// func sortList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}

// 	start := &ListNode{}
// 	start.Next = head

// 	end := head
// 	cur := end.Next
// 	sorted := make([]*ListNode, 0)
// 	sorted = append(sorted, head)
// 	for cur != nil {
// 		if end.Val <= cur.Val {
// 			sorted = append(sorted, cur)
// 			end = cur
// 			cur = end.Next
// 		} else if start.Next.Val >= cur.Val {
// 			sorted = append([]*ListNode{cur}, sorted...)
// 			end.Next = cur.Next
// 			cur.Next = start.Next
// 			start.Next = cur
// 			cur = end.Next

// 		} else {
// 			index := bs(0, len(sorted)-1, cur, sorted)
// 			p := sorted[index]
// 			sortedPre := []*ListNode{}
// 			sortedPre = append(sortedPre, sorted[:index+1]...)
// 			sortedPre = append(sortedPre, cur)
// 			sortedPre = append(sortedPre, sorted[index+1:]...)
// 			sorted = sortedPre
// 			end.Next = cur.Next
// 			cur.Next = p.Next
// 			p.Next = cur
// 			cur = end.Next
// 		}
// 	}
// 	return start.Next
// }

// func bs(l, r int, cur *ListNode, sorted []*ListNode) int {
// 	if l == r {
// 		return l
// 	}

// 	mid := (l + r) >> 1

// 	if sorted[mid].Val == cur.Val {
// 		return mid
// 	}

// 	if sorted[mid].Val > cur.Val {
// 		return bs(l, mid, cur, sorted)
// 	} else {
// 		if sorted[mid].Next.Val >= cur.Val {
// 			return mid
// 		}
// 		return bs(mid+1, r, cur, sorted)
// 	}
// }

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	right := slow.Next
	slow.Next = nil
	left := sortList(head)
	right = sortList(right)
	return mergeSortedList(left, right)
}

func mergeSortedList(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	start := &ListNode{}
	p := start
	for left != nil || right != nil {
		if left != nil && right != nil {
			if left.Val < right.Val {
				p.Next = left
				p = p.Next
				left = left.Next
				continue
			} else if left.Val == right.Val {
				p.Next = left
				p = p.Next
				left = left.Next

				p.Next = right
				p = p.Next
				right = right.Next
				continue
			} else {
				p.Next = right
				p = p.Next
				right = right.Next
				continue
			}
		} else if left != nil {
			p.Next = left
			p = p.Next
			left = left.Next
			continue
		} else {
			p.Next = right
			p = p.Next
			right = right.Next
			continue
		}
	}
	return start.Next
}
