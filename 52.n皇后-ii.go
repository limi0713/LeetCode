/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 *
 * https://leetcode-cn.com/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (77.73%)
 * Likes:    108
 * Dislikes: 0
 * Total Accepted:    20.1K
 * Total Submissions: 25.8K
 * Testcase Example:  '4'
 *
 * n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 * 
 * 
 * 
 * 上图为 8 皇后问题的一种解法。
 * 
 * 给定一个整数 n，返回 n 皇后不同的解决方案的数量。
 * 
 * 示例:
 * 
 * 输入: 4
 * 输出: 2
 * 解释: 4 皇后问题存在如下两个不同的解法。
 * [
 * [".Q..",  // 解法 1
 * "...Q",
 * "Q...",
 * "..Q."],
 * 
 * ["..Q.",  // 解法 2
 * "Q...",
 * "...Q",
 * ".Q.."]
 * ]
 * 
 * 
 */

// @lc code=start
func totalNQueens(n int) int {
	var ans int = 0
	flag := make([][]int, 0)
	for i := 0; i < n; i ++{
		f := make([]int, n)
		flag = append(flag, f)
	}

	dfs(&ans, n, 0, flag)

	return ans
}

func dfs(ans *int, n, row int, flag [][]int) {
	if row == n {
		*ans = *ans + 1
		return
	}
	
	if row > n {
		return
	}

	for i := 0; i < n; i ++ {
		if flag[row][i] == 0 {
			setFlag(flag, n, row, i)

			dfs(ans, n, row + 1, flag)

			resetFlag(flag, n, row, i)
		}
	}

}

func setFlag(flag [][]int, n, row, col int){
	for i := 0; i < n; i ++ {
		flag[row][i] += 1
		flag[i][col] += 1

		if row - i >= 0 && col - i >= 0{
			flag[row-i][col-i] += 1
		}

		if row - i >= 0 && col + i < n {
			flag[row-i][col+i] += 1
		}

		if row + i < n && col - i >= 0{
			flag[row+i][col-i] += 1
		}

		if row + i < n && col + i < n {
			flag[row+i][col+i] += 1
		}
	}
}

func resetFlag(flag [][]int, n, row, col int){
	for i := 0; i < n; i ++ {
		flag[row][i] -= 1
		flag[i][col] -= 1

		if row - i >= 0 && col - i >= 0{
			flag[row-i][col-i] -= 1
		}

		if row - i >= 0 && col + i < n {
			flag[row-i][col+i] -= 1
		}

		if row + i < n && col - i >= 0{
			flag[row+i][col-i] -= 1
		}

		if row + i < n && col + i < n {
			flag[row+i][col+i] -= 1
		}
	}
}
// @lc code=end

