/*
 * @lc app=leetcode.cn id=938 lang=golang
 *
 * [938] 二叉搜索树的范围和
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

func cacu(root *TreeNode, L, R int) int {
	if root == nil {
		return 0
	}
	
	l := 0
	r := 0

	l = cacu(root.Left,L,R)
	r = cacu(root.Right,L,R)

	if root.Val>=L && root.Val <= R {
		l = cacu(root.Left,L,R)
		r = cacu(root.Right,L,R)
		return l+root.Val+r
	}else if root.Val < L {
		r = cacu(root.Right,L,R)
		return r
	}else if root.Val > R {
		l = cacu(root.Left,L,R)
		return l
	}
	return 0
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	res := 0
	res = cacu(root,L,R)
	return res
}

// @lc code=end

