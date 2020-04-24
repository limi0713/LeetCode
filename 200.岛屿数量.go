/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 *
 * https://leetcode-cn.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (48.00%)
 * Likes:    500
 * Dislikes: 0
 * Total Accepted:    87.3K
 * Total Submissions: 179K
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
 * 
 * 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
 * 
 * 此外，你可以假设该网格的四条边均被水包围。
 * 
 * 示例 1:
 * 
 * 输入:
 * 11110
 * 11010
 * 11000
 * 00000
 * 输出: 1
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * 11000
 * 11000
 * 00100
 * 00011
 * 输出: 3
 * 解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
 * 
 * 
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	ans := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				ans ++
				dfs(&grid,i,j)
			}
		}
	}
	return ans
}

func dfs(grid *[][]byte, row, col int) {
	(*grid)[row][col] = '2'

	if row - 1 >= 0 && (*grid)[row-1][col] == '1' {
		dfs(grid,row-1,col)
	}

	if row + 1 < len(*grid) && (*grid)[row+1][col] == '1' {
		dfs(grid,row+1,col)
	}

	if col - 1 >= 0 && (*grid)[row][col-1] == '1' {
		dfs(grid,row,col-1)
	}

	if col + 1 < len((*grid)[row]) && (*grid)[row][col+1] == '1' {
		dfs(grid,row,col+1)
	}
}
// @lc code=end

