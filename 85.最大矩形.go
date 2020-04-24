/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 *
 * https://leetcode-cn.com/problems/maximal-rectangle/description/
 *
 * algorithms
 * Hard (45.36%)
 * Likes:    387
 * Dislikes: 0
 * Total Accepted:    24.6K
 * Total Submissions: 54.2K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * 给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
 * 
 * 示例:
 * 
 * 输入:
 * [
 * ⁠ ["1","0","1","0","0"],
 * ⁠ ["1","0","1","1","1"],
 * ⁠ ["1","1","1","1","1"],
 * ⁠ ["1","0","0","1","0"]
 * ]
 * 输出: 6
 * 
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {

	heights := make([][]int, 0,len(matrix))
	for i := range matrix {
		h := make([]int,len(matrix[i]))
		heights = append(heights,h)
	}

	for i := range matrix {
		for j := range matrix[i] {
			heights[i][j] = int(matrix[i][j] - '0')

			if matrix[i][j] == '1' && i - 1 >= 0 {
				heights[i][j] += heights[i-1][j]
			}
		}
	}

	maxAns := 0

	for i := range heights {
		maxAns = max(maxAns,stack(heights[i]))
	}

	return maxAns
}

func stack(heights []int) int {
	ls := list.New()

	ls.PushBack(-1)

	maxAns := 0
	for i := range heights {
		for ls.Len() > 1 {
			index := ls.Back().Value.(int)
			if heights[i] <= heights[index] {
				ls.Remove(ls.Back())
				indexBefore := ls.Back().Value.(int)
				maxAns = max(maxAns,heights[index] * (i-indexBefore - 1))
			}else{
				break
			}
		}

		ls.PushBack(i)
	}

	for ls.Len() > 1 {
		index := ls.Back().Value.(int)
		ls.Remove(ls.Back())
		indexBefore := ls.Back().Value.(int)
		maxAns = max(maxAns,heights[index]*(len(heights) - indexBefore - 1))
	}

	return maxAns
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

