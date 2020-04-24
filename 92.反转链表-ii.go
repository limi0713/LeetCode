/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 *
 * https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (49.68%)
 * Likes:    350
 * Dislikes: 0
 * Total Accepted:    45.2K
 * Total Submissions: 90.9K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
 * 
 * 说明:
 * 1 ≤ m ≤ n ≤ 链表长度。
 * 
 * 示例:
 * 
 * 输入: 1->2->3->4->5->NULL, m = 2, n = 4
 * 输出: 1->4->3->2->5->NULL
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
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	temp := &ListNode{
		Next:head,
	}

	var insertHere *ListNode
	p1 := temp
	index := 0
	for p1 != nil && p1.Next != nil {
		if index == m - 1 {
			insertHere = p1
		}

		next := p1.Next
		
		if index >= m && index < n {
			p1.Next = next.Next
			next.Next = insertHere.Next
			insertHere.Next = next
			index ++
			continue
		}

		p1 = p1.Next
		index ++

	}

	return temp.Next

}
// @lc code=end

