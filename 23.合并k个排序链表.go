/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 *
 * https://leetcode-cn.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (49.65%)
 * Likes:    557
 * Dislikes: 0
 * Total Accepted:    91.5K
 * Total Submissions: 184.3K
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
 * 
 * 示例:
 * 
 * 输入:
 * [
 * 1->4->5,
 * 1->3->4,
 * 2->6
 * ]
 * 输出: 1->1->2->3->4->4->5->6
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoList(listA,listB *ListNode)*ListNode{
	if listA == nil {
		return listB
	}

	if listB == nil {
		return listA
	}

	var ans *ListNode
	p1 ,p2 := listA,listB
	if p1.Val < p2.Val{
		ans = p1
		p1 = p1.Next
	}else {
		ans = p2
		p2 = p2.Next
	}

	p3 := ans

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val{
			p3.Next = p1
			p1 = p1.Next
		}else {
			p3.Next = p2
			p2 = p2.Next
		}

		p3 = p3.Next
	}

	for p1 != nil {
		p3.Next = p1
		p1 = p1.Next
		p3 = p3.Next
	}

	for p2 != nil {
		p3.Next = p2
		p2 = p2.Next
		p3 = p3.Next

	}

	return ans
}

func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) <= 0 {
		return nil
	}

	if len(lists) == 1{
		return lists[0]
	}

	return mergeTwoList(lists[0],mergeKLists(lists[1:]))
}
// @lc code=end

