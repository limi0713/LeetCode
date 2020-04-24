/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (62.39%)
 * Likes:    359
 * Dislikes: 0
 * Total Accepted:    27.4K
 * Total Submissions: 43.9K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
 * 
 * 示例:
 * 
 * 输入: 3
 * 输出:
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * 解释:
 * 以上的输出对应以下 5 种不同结构的二叉搜索树：
 * 
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
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
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return generateRange(1, n)
}

func generateRange(left, right int) []*TreeNode {
	if left > right {
		return []*TreeNode{nil}
	}

	res := make([]*TreeNode,0)

	for i:=left;i<=right;i++{
		leftITrees,rightITrees := generateRange(left,i-1),generateRange(i+1,right)

		for j := range leftITrees {
			for k := range rightITrees {
				root := &TreeNode{
					Val:i,
					Left:leftITrees[j],
					Right:rightITrees[k],
				}

				res = append(res,root)
			}
		}
	}

	return res
}

// @lc code=end

