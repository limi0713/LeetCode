/*
 * @lc app=leetcode.cn id=958 lang=golang
 *
 * [958] 二叉树的完全性检验
 *
 * https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree/description/
 *
 * algorithms
 * Medium (46.99%)
 * Likes:    34
 * Dislikes: 0
 * Total Accepted:    4K
 * Total Submissions: 8.5K
 * Testcase Example:  '[1,2,3,4,5,6]'
 *
 * 给定一个二叉树，确定它是否是一个完全二叉树。
 * 
 * 百度百科中对完全二叉树的定义如下：
 * 
 * 若设二叉树的深度为 h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h
 * 层所有的结点都连续集中在最左边，这就是完全二叉树。（注：第 h 层可能包含 1~ 2^h 个节点。）
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：[1,2,3,4,5,6]
 * 输出：true
 * 解释：最后一层前的每一层都是满的（即，结点值为 {1} 和 {2,3} 的两层），且最后一层中的所有结点（{4,5,6}）都尽可能地向左。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 输入：[1,2,3,4,5,null,7]
 * 输出：false
 * 解释：值为 7 的结点没有尽可能靠向左侧。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中将会有 1 到 100 个结点。
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



func dfs(node *TreeNode)(bool,int,int){
	if node == nil {
		return true,0,0
	}
	
	leftSonTree, leftMax, leftMin := dfs(node.Left)
	rightSonTree, rightMax, rightMin := dfs(node.Right)

	if leftMin < rightMax {
		return false, 0,0
	}

	if leftMax != rightMax && leftMax - 1 != rightMax {
		return false, 0,0
	}

	if leftMax != rightMax && leftMax - 1 != rightMax {
		return false, 0,0
	}

	if rightMin + 1 < rightMax || leftMin + 1 < leftMax || rightMin + 1 < leftMax{
		return false, 0,0
	}


	max, _ := comp(leftMax,rightMax)
	_, min := comp(leftMin, rightMin)

	return leftSonTree && rightSonTree, max+1, min+1
}

func comp(a, b int) (int, int) {
	if a < b {
		return b, a
	}
	return a,b
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftSonTree, leftMax, leftMin := dfs(root.Left)
	rightSonTree, rightMax, rightMin := dfs(root.Right)

	if leftMin < rightMax {
		return false
	}

	if leftMax != rightMax && leftMax - 1 != rightMax {
		return false
	}

	if rightMin + 1 < rightMax || leftMin + 1 < leftMax || rightMin + 1 < leftMax{
		return false
	}

	return leftSonTree && rightSonTree
}
// @lc code=end

