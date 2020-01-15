/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 *
 * https://leetcode-cn.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (49.21%)
 * Likes:    557
 * Dislikes: 0
 * Total Accepted:    81.1K
 * Total Submissions: 164.2K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给定一个二叉树，检查它是否是镜像对称的。
 * 
 * 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 * 
 * 
 * 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 * 
 * 
 * 说明:
 * 
 * 如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
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

var (
	midList []int
)

func midDP(root *TreeNode){
	if root == nil{
		return
	}

	midDP(root.Left)
	midList = append(midList,root.Val)
	midDP(root.Right)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	midList = make([]int, 0)
	midDP(root)

	var l int = 0
	var r int = len(midList) - 1
	if r == 0{
		return true
	}

	for l <= r {

		if midList[l] != midList[r]{
			return false
		}

		l++
		r--
	}

	return true
}
// @lc code=end

