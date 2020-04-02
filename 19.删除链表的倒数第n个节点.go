/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 *
 * https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (38.21%)
 * Likes:    761
 * Dislikes: 0
 * Total Accepted:    146.7K
 * Total Submissions: 384K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
 * 
 * 示例：
 * 
 * 给定一个链表: 1->2->3->4->5, 和 n = 2.
 * 
 * 当删除了倒数第二个节点后，链表变为 1->2->3->5.
 * 
 * 
 * 说明：
 * 
 * 给定的 n 保证是有效的。
 * 
 * 进阶：
 * 
 * 你能尝试使用一趟扫描实现吗？
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

func dfs(node *ListNode,n int)int{

	if node.Next == nil {
		return 1
	}

	indexNext := dfs(node.Next,n)

	if indexNext == n {
		node.Next = node.Next.Next
	}


	return 1 + indexNext 
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return nil
	}
	
	// if n == 1 && head.Next == nil {
	// 	return nil
	// }

	temp := &ListNode{}
	temp.Next = head


	dfs(temp,n)

	return temp.Next
}
// @lc code=end

