/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (67.85%)
 * Likes:    278
 * Dislikes: 0
 * Total Accepted:    60.5K
 * Total Submissions: 88.9K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的中序 遍历。
 *
 * 示例:
 *
 * 输入: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * 输出: [1,3,2]
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
递归
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var rel []int
	if root.Left != nil {
		rel = inorderTraversal(root.Left)
	}
	rel = append(rel, root.Val)
	var rer []int
	if root.Right != nil {
		rer = inorderTraversal(root.Right)
	}
	for _, v := range rer {
		rel = append(rel, v)
	}
	return rel
}
*/

//stack
type stack []*TreeNode

func (s stack) Push(v *TreeNode) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, *TreeNode) {
	l := len(s)
	if l == 0 {
		return s, nil
	}
	return s[:l-1], s[l-1]
}

//迭代
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var re []int
	var s stack
	s = s.Push(root)
	for len(s) > 0 {
		var r *TreeNode
		s, r = s.Pop()
		for r != nil {
			if r.Left != nil {
				t := r.Left
				r.Left = nil
				s = s.Push(r)
				r = t
			} else {
				re = append(re, r.Val)
				r = r.Right
			}
		}
	}
	return re
}
