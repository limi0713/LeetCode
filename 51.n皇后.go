/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N皇后
 *
 * https://leetcode-cn.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (68.73%)
 * Likes:    371
 * Dislikes: 0
 * Total Accepted:    34.1K
 * Total Submissions: 49.6K
 * Testcase Example:  '4'
 *
 * n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 * 
 * 
 * 
 * 上图为 8 皇后问题的一种解法。
 * 
 * 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
 * 
 * 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
 * 
 * 示例:
 * 
 * 输入: 4
 * 输出: [
 * ⁠[".Q..",  // 解法 1
 * ⁠ "...Q",
 * ⁠ "Q...",
 * ⁠ "..Q."],
 * 
 * ⁠["..Q.",  // 解法 2
 * ⁠ "Q...",
 * ⁠ "...Q",
 * ⁠ ".Q.."]
 * ]
 * 解释: 4 皇后问题存在两个不同的解法。
 * 
 * 
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	strs := make([]string,0)
	for i :=0; i < n; i ++ {
		str :=getPoint(i) + "Q" + getPoint(n - i - 1)
		strs = append(strs, str)
	}

	ans := make([][]string, 0)

	flag := make([][]int,0)
	for i := 0; i < n; i ++ {
		f := make([]int,n)
		flag = append(flag,f)
	}

	dfs(&ans, n, 0, strs, []string{}, flag)

	return ans
}

func dfs(ans *[][]string, n, row int, strs, cur []string, flag [][]int){
	if row == n && len(cur) == n {
		*ans = append(*ans, cur)
		return
	}

	if row >= n {
		return
	}

	for i := 0; i < len(strs); i ++ {
		if flag[row][i] == 0 {
			setFlag(flag,n,row,i)

			temp := make([]string,0)
			temp = append(temp,cur...)
			temp = append(temp, strs[i])

			dfs(ans, n, row + 1, strs, temp, flag)

			resetFlag(flag,n,row,i)
		}
	}

}

func setFlag(flag [][]int, n, row, column int){

	for i:= 0; i < n; i ++ {
		flag[row][i] += 1 
		flag[i][column] += 1 

		if row - i >= 0 && column - i >= 0 {
			flag[row-i][column-i] += 1 
		} 

		if row - i >= 0 && column + i < n {
			flag[row-i][column+i] += 1 
		} 

		if row + i < n && column - i >= 0 {
			flag[row+i][column-i] += 1 
		} 

		if row + i < n && column + i < n {
			flag[row+i][column+i] += 1 
		} 

	}
}

func resetFlag(flag [][]int, n, row, column int){

	for i:= 0; i < n; i ++ {
		flag[row][i] -= 1 
		flag[i][column] -= 1

		if row - i >= 0 && column - i >= 0 {
			flag[row-i][column-i] -= 1
		} 

		if row - i >= 0 && column + i < n {
			flag[row-i][column+i] -= 1
		} 

		if row + i < n && column - i >= 0 {
			flag[row+i][column-i] -= 1
		} 

		if row + i < n && column + i < n {
			flag[row+i][column+i] -= 1
		} 

	}
}


func getPoint(n int) string {
	res := ""
	for i := 0; i < n; i ++ {
		res += "."
	}

	return res
}
// @lc code=end

