/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 *
 * https://leetcode.cn/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (56.98%)
 * Likes:    2034
 * Dislikes: 0
 * Total Accepted:    492.6K
 * Total Submissions: 864.1K
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 给你一个链表数组，每个链表都已经按升序排列。
 *
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 *
 *
 *
 * 示例 1：
 *
 * 输入：lists = [[1,4,5],[1,3,4],[2,6]]
 * 输出：[1,1,2,3,4,4,5,6]
 * 解释：链表数组如下：
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * 将它们合并到一个有序链表中得到。
 * 1->1->2->3->4->4->5->6
 *
 *
 * 示例 2：
 *
 * 输入：lists = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 * 输入：lists = [[]]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] 按 升序 排列
 * lists[i].length 的总和不超过 10^4
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

func merge(p1, p2 *ListNode) *ListNode {
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

		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) {
				mRet = append(mRet, merge(lists[i], lists[i+1]))
			} else {
				mRet = append(mRet, lists[i])
			}
		}
		lists = append(make([]*ListNode, 0), mRet...)
	}
}

// @lc code=end

