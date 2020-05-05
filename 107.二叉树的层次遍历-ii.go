/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (65.09%)
 * Likes:    227
 * Dislikes: 0
 * Total Accepted:    54.9K
 * Total Submissions: 84.3K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
 * 
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 返回其自底向上的层次遍历为：
 * 
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	levIndex := 0
	levMap := make(map[int][]*TreeNode)
	levMap[0] = []*TreeNode{root}
	// ans := make([][]int,0)

	for len(levMap[levIndex]) > 0 {
		temp := make([]*TreeNode,0,2*len(levMap[levIndex]))
		// a := make([]int,0,len(levMap[levIndex]))
		for _, v := range levMap[levIndex] {
			// a = append(a,v.Val)

			if v.Left != nil {
				temp = append(temp,v.Left)
			}

			if v.Right != nil {
				temp = append(temp,v.Right)
			}
		}

		// ans = append([][]int{a},ans...)

		levIndex ++
		if len(temp) > 0 {
			levMap[levIndex] = temp
		}
	}

	ans := make([][]int,levIndex)
	for k,v:=range levMap{
		ans[levIndex-k-1] = make([]int,0,len(v))
		for i := range v {
			ans[levIndex-k-1] = append(ans[levIndex-k-1], v[i].Val)
		}
	}

	return ans
}
// @lc code=end

