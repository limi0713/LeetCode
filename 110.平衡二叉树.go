/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 *
 * https://leetcode-cn.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (51.49%)
 * Likes:    302
 * Dislikes: 0
 * Total Accepted:    71.2K
 * Total Submissions: 138.2K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，判断它是否是高度平衡的二叉树。
 * 
 * 本题中，一棵高度平衡二叉树定义为：
 * 
 * 
 * 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
 * 
 * 
 * 示例 1:
 * 
 * 给定二叉树 [3,9,20,null,null,15,7]
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 返回 true 。
 * 
 * 示例 2:
 * 
 * 给定二叉树 [1,2,2,3,3,null,null,4,4]
 * 
 * ⁠      1
 * ⁠     / \
 * ⁠    2   2
 * ⁠   / \
 * ⁠  3   3
 * ⁠ / \
 * ⁠4   4
 * 
 * 
 * 返回 false 。
 * 
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
func isBalanced(root *TreeNode) bool {
	ans,_ := DFS(root)
	return ans
}

func DFS(node *TreeNode) (bool,int) {
	if node == nil {
		return true,0
	}

	flag,leftLev := DFS(node.Left)
	if !flag {
		return false,0
	}

	flag,rightLev := DFS(node.Right)
	if !flag {
		return false,0
	}

	max,min := comp(leftLev,rightLev)
	if max-min > 1 {
		return false,0
	}

	return true,max+1

}

func comp(a, b int) (int,int) {
	if a > b {
		return a, b
	}

	return b,a
}
// @lc code=end

