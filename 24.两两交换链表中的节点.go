/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (65.01%)
 * Likes:    455
 * Dislikes: 0
 * Total Accepted:    88.2K
 * Total Submissions: 135.7K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 * 
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 * 
 * 
 * 
 * 示例:
 * 
 * 给定 1->2->3->4, 你应该返回 2->1->4->3.
 * 
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
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tempNode := &ListNode{}
	tempNode.Next = head

	p1 := tempNode

	for p1.Next != nil && p1.Next.Next != nil {

		p2 := p1.Next
		p3 := p2.Next
		p4 := p3.Next

		p1.Next = p3
		p3.Next = p2
		p2.Next = p4

		p1 = p2
	}

	return tempNode.Next

}
// @lc code=end

