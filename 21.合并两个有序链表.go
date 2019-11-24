/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var res *ListNode
	if l1.Val <= l2.Val {
		res = l1
		l1 = l1.Next
	}else {
		res = l2
		l2 = l2.Next
	}
	p1 := res

	for l1!=nil &&l2!=nil {
		if l1.Val <= l2.Val {
			p1.Next = l1
			l1 = l1.Next
		}else {
			p1.Next = l2
			l2 = l2.Next
		}
		p1=p1.Next
	}

	for l1 != nil {
		p1.Next = l1
		l1 = l1.Next
		p1=p1.Next
	}

	for l2 != nil {
		p1.Next = l2
		l2 = l2.Next
		p1=p1.Next
	}

	return res
}
// @lc code=end

