/*
 * @lc app=leetcode.cn id=1161 lang=golang
 *
 * [1161] 最大层内元素和
 *
 * https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree/description/
 *
 * algorithms
 * Medium (62.87%)
 * Likes:    60
 * Dislikes: 0
 * Total Accepted:    14.7K
 * Total Submissions: 23.1K
 * Testcase Example:  '[1,7,0,7,-8,null,null]'
 *
 * 给你一个二叉树的根节点 root。设根节点位于二叉树的第 1 层，而根节点的子节点位于第 2 层，依此类推。
 *
 * 请返回层内元素之和 最大 的那几层（可能只有一层）的层号，并返回其中 最小 的那个。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [1,7,0,7,-8,null,null]
 * 输出：2
 * 解释：
 * 第 1 层各元素之和为 1，
 * 第 2 层各元素之和为 7 + 0 = 7，
 * 第 3 层各元素之和为 7 + -8 = -1，
 * 所以我们返回第 2 层的层号，它的层内元素之和最大。
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [989,null,10250,98693,-89388,null,null,null,-32127]
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数在 [1, 10^4]范围内
 * -10^5 <= Node.val <= 10^5
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
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	currentLevelIndex := 0
	currentLevelNodes := make([]*TreeNode, 0)
	nextLevelNodes := make([]*TreeNode, 0)
	nextLevelNodes = append(nextLevelNodes, root)
	maxSumLevel := 1
	maxSum := root.Val

	for len(nextLevelNodes) > 0 {
		currentLevelNodes = append([]*TreeNode{}, nextLevelNodes...)
		nextLevelNodes = make([]*TreeNode, 0)
		currentSum := 0
		currentLevelIndex++

		for _, node := range currentLevelNodes {
			if node == nil {
				continue
			}

			currentSum += node.Val
			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}
			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}

		if currentSum > maxSum {
			maxSum = currentSum
			maxSumLevel = currentLevelIndex
		}
	}

	return maxSumLevel
}

// @lc code=end

