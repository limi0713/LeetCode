/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 *
 * https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (72.99%)
 * Likes:    517
 * Dislikes: 0
 * Total Accepted:    166.3K
 * Total Submissions: 227.7K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最大深度。
 * 
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例：
 * 给定二叉树 [3,9,20,null,null,15,7]，
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 返回它的最大深度 3 。
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	levIndex := 0
	levMap := make(map[int][]*TreeNode)

	levMap[levIndex] = []*TreeNode{root}
	for len(levMap[levIndex]) > 0 {
		temp := make([]*TreeNode,0)

		for _,v := range levMap[levIndex] {
			if v.Left != nil {
				temp = append(temp, v.Left)
			}

			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}

		delete(levMap,levIndex)
		levIndex ++
		levMap[levIndex] = temp
	}

	return levIndex
}
// @lc code=end

