/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (48.44%)
 * Likes:    238
 * Dislikes: 0
 * Total Accepted:    63.7K
 * Total Submissions: 131.1K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
 * 
 * 示例 1:
 * 
 * 输入: 1->1->2
 * 输出: 1->2
 * 
 * 
 * 示例 2:
 * 
 * 输入: 1->1->2->3->3
 * 输出: 1->2->3
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
    if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	p1 := head
	p2 := head.Next
	for p2 != nil{
		if p2.Val == p1.Val {
			p1.Next = p2.Next
			p2 = p2.Next
		}else{
			p1 = p2
			p2 = p2.Next
		}
	}

	return head
}
// @lc code=end

