/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 *
 * https://leetcode-cn.com/problems/word-search/description/
 *
 * algorithms
 * Medium (41.22%)
 * Likes:    380
 * Dislikes: 0
 * Total Accepted:    51.4K
 * Total Submissions: 124.7K
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
 * 
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
 * 
 * 
 * 
 * 示例:
 * 
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 * 
 * 给定 word = "ABCCED", 返回 true
 * 给定 word = "SEE", 返回 true
 * 给定 word = "ABCB", 返回 false
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * board 和 word 中只包含大写和小写英文字母。
 * 1 <= board.length <= 200
 * 1 <= board[i].length <= 200
 * 1 <= word.length <= 10^3
 * 
 * 
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return len(word) == 0
	}

	flag := make([][]bool, 0)
	for i := 0; i < len(board); i ++ {
		f := make([]bool, len(board[i]))
		flag = append(flag, f)
	}

	for i := 0; i < len(board); i ++ {
		for j := 0; j < len(board[i]); j ++ {
			if board[i][j] == word[0] {
				if backtrack(&board, i, j, word, 0, &flag) {	// board[i][j] == word[0]
					return true
				}
			}
		}
	}

	return false

}

func backtrack(board *[][]byte, row, col int, word string, index int, flag *[][]bool) bool {
	
	(*flag)[row][col] = true

	if index == len(word) - 1 {
		return true
	}

	if row - 1 >= 0 && (*flag)[row-1][col] == false && (*board)[row-1][col] == word[index+1]  {
		if backtrack(board, row-1, col, word, index + 1, flag) {
			return true
		}
	}

	if row + 1 < len(*board) && (*flag)[row+1][col] == false && (*board)[row+1][col] == word[index+1] {
		if backtrack(board, row+1, col, word, index+1, flag) {
			return true
		}
	}

	if col - 1 >= 0 && (*flag)[row][col-1] == false  && (*board)[row][col-1] == word[index+1] {
		if backtrack(board, row, col-1, word, index+1, flag) {
			return true
		}
	}

	if col + 1 < len((*board)[row]) && (*flag)[row][col+1] == false && (*board)[row][col+1] == word[index+1] {
		if backtrack(board, row, col+1, word, index+1, flag) {
			return true
		}
	}

	(*flag)[row][col] = false

	return false

}
// @lc code=end

