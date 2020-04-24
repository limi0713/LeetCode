/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 *
 * https://leetcode-cn.com/problems/partition-list/description/
 *
 * algorithms
 * Medium (57.03%)
 * Likes:    185
 * Dislikes: 0
 * Total Accepted:    33.4K
 * Total Submissions: 58.4K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * 给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
 * 
 * 你应当保留两个分区中每个节点的初始相对位置。
 * 
 * 示例:
 * 
 * 输入: head = 1->4->3->2->5->2, x = 3
 * 输出: 1->2->2->4->3->5
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
func partition(head *ListNode, x int) *ListNode {
	temp := &ListNode{}
	temp.Next = head
	curNode, validNode := temp, temp

	for curNode != nil && curNode.Next != nil {
		nextNode := curNode.Next

		if nextNode.Val >= x {
			curNode = curNode.Next
			continue
		}

		if curNode == validNode {
			validNode = validNode.Next
			curNode = nextNode
			continue
		}

		curNode.Next = nextNode.Next
		nextNode.Next = validNode.Next
		validNode.Next = nextNode
		validNode = validNode.Next
	}

	return temp.Next
}
// @lc code=end

