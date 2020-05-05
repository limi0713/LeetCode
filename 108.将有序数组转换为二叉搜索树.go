/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
 *
 * https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (70.66%)
 * Likes:    399
 * Dislikes: 0
 * Total Accepted:    64.8K
 * Total Submissions: 91.6K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
 * 
 * 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
 * 
 * 示例:
 * 
 * 给定有序数组: [-10,-3,0,5,9],
 * 
 * 一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
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
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
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

	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root 
}
// @lc code=end

