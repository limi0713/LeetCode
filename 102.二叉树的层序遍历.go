/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (61.68%)
 * Likes:    460
 * Dislikes: 0
 * Total Accepted:    110.9K
 * Total Submissions: 179.7K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
 * 
 * 
 * 
 * 示例：
 * 二叉树：[3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 返回其层次遍历结果：
 * 
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := make([][]int,0)

	levMap := make(map[int][]*TreeNode)

	levIndex := 0

	levMap[levIndex] = []*TreeNode{root}

	for len(levMap[levIndex]) > 0 {
		temp := make([]*TreeNode,0)

		a := make([]int,0)
		for i := range levMap[levIndex]{
			a = append(a,levMap[levIndex][i].Val)
			if levMap[levIndex][i].Left != nil {
				temp = append(temp,levMap[levIndex][i].Left)
			}

			if levMap[levIndex][i].Right != nil {
				temp = append(temp,levMap[levIndex][i].Right)
			}
		}

		ans = append(ans,a)

		delete(levMap,levIndex)
		levMap[levIndex+1] = temp
		levIndex ++
	}

	return ans
}
// @lc code=end

