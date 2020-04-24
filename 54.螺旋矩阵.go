/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode-cn.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (39.28%)
 * Likes:    340
 * Dislikes: 0
 * Total Accepted:    51.4K
 * Total Submissions: 130.6K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
 * 
 * 示例 1:
 * 
 * 输入:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 * 输出: [1,2,3,6,9,8,7,4,5]
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * [
 * ⁠ [1, 2, 3, 4],
 * ⁠ [5, 6, 7, 8],
 * ⁠ [9,10,11,12]
 * ]
 * 输出: [1,2,3,4,8,12,11,10,9,5,6,7]
 * 
 * 
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {

	if len(matrix) <= 0 {
		return []int{}
	}

	ans := make([]int, 0)

	m, n := len(matrix), len(matrix[0])

	ring := (min(m, n) - 1)/2 + 1

	for i :=0 ; i < ring; i ++ {
		
		for j := i; j <= n - i - 1; j ++ {
			ans = append(ans, matrix[i][j])
		}

		for j := i+1; j <= m - i - 1;j ++ {
			ans = append(ans, matrix[j][n - 1 -i])
		}

		if m -1 -i > i {
			for j := n - 2 - i ; j >= i; j -- {
				ans = append(ans, matrix[m-1-i][j])
			}
		}

		if i < n - 1 - i {
			for j := m - 2 - i ; j >= i+1; j -- {
				ans = append(ans, matrix[j][i])
			}
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

