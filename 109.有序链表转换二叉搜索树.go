/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 *
 * https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/description/
 *
 * algorithms
 * Medium (71.88%)
 * Likes:    197
 * Dislikes: 0
 * Total Accepted:    26K
 * Total Submissions: 36.2K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
 * 
 * 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
 * 
 * 示例:
 * 
 * 给定的有序链表： [-10, -3, 0, 5, 9],
 * 
 * 一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
 * 
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	nums := make([]int,0)
	listToArr(head,&nums)
	return arrToBST(nums)
}

func listToArr(node *ListNode,nums *[]int){
	for node != nil {
		*nums = append(*nums,node.Val)
		node = node.Next
	}
}

func arrToBST(nums []int) * TreeNode{
	if len(nums) <= 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{
			Val:nums[0],
		}
	}

	if len(nums) == 2 {
		return &TreeNode{
			Val:nums[1],
			Left:&TreeNode{
				Val:nums[0],
			},
		}
	}

	if len(nums) == 3 {
		return &TreeNode{
			Val:nums[1],
			Left:&TreeNode{
				Val:nums[0],
			},
			Right:&TreeNode{
				Val:nums[2],
			},
		}
	}

	mid := len(nums) >> 1

	root := &TreeNode{
		Val:nums[mid],
	}

	root.Left = arrToBST(nums[:mid])
	root.Right = arrToBST(nums[mid+1:])
	return root 
}

// @lc code=end

