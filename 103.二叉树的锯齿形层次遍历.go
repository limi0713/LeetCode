/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (54.29%)
 * Likes:    184
 * Dislikes: 0
 * Total Accepted:    45.2K
 * Total Submissions: 83.3K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 * 
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 返回锯齿形层次遍历如下：
 * 
 * [
 * ⁠ [3],
 * ⁠ [20,9],
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := make([][]int,0)
	levMap, levIndex := make(map[int][]*TreeNode), 0

	levMap[levIndex] = []*TreeNode{root}
	for len(levMap[levIndex]) > 0 {
		a := make([]int,len(levMap[levIndex]))

		temp := make([]*TreeNode,0)
		
		for i :=0;i<len(levMap[levIndex]);i++{
			if levIndex % 2 == 0 {
				a[i] = levMap[levIndex][i].Val
			}else {
				a[len(a)-1-i] = levMap[levIndex][i].Val
			}
			if levMap[levIndex][i].Left != nil {
				temp = append(temp, levMap[levIndex][i].Left)
			}

			if levMap[levIndex][i].Right != nil {
				temp = append(temp, levMap[levIndex][i].Right)
			}

		}

		delete(levMap,levIndex)

		levMap[levIndex+1] = temp
		levIndex ++
		ans = append(ans,a)
	}

	return ans
}
// @lc code=end

