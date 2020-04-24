/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (47.22%)
 * Likes:    256
 * Dislikes: 0
 * Total Accepted:    41.9K
 * Total Submissions: 88.7K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
 * 
 * 示例 1:
 * 
 * 输入: 1->2->3->3->4->4->5
 * 输出: 1->2->5
 * 
 * 
 * 示例 2:
 * 
 * 输入: 1->1->1->2->3
 * 输出: 2->3
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
func deleteDuplicates(head *ListNode) *ListNode {

	tempNode := &ListNode{}
	tempNode.Next = head
	p1 := tempNode

	for p1 != nil && p1.Next != nil {
		cur := p1.Next

		deleteFlag := false

		for cur != nil && cur.Next != nil {
			if cur.Next.Val == cur.Val {
				deleteFlag = true
				cur.Next = cur.Next.Next
			}else {
				break
			}
		}

		if deleteFlag {
			p1.Next = cur.Next
		}else {
			p1 = p1.Next
		}
	}

	return tempNode.Next
}
// @lc code=end

