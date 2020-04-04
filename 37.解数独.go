/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 *
 * https://leetcode-cn.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (60.66%)
 * Likes:    367
 * Dislikes: 0
 * Total Accepted:    23.4K
 * Total Submissions: 38.5K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * 编写一个程序，通过已填充的空格来解决数独问题。
 * 
 * 一个数独的解法需遵循如下规则：
 * 
 * 
 * 数字 1-9 在每一行只能出现一次。
 * 数字 1-9 在每一列只能出现一次。
 * 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
 * 
 * 
 * 空白格用 '.' 表示。
 * 
 * 
 * 
 * 一个数独。
 * 
 * 
 * 
 * 答案被标成红色。
 * 
 * Note:
 * 
 * 
 * 给定的数独序列只包含数字 1-9 和字符 '.' 。
 * 你可以假设给定的数独只有唯一解。
 * 给定数独永远是 9x9 形式的。
 * 
 * 
 */

// @lc code=start
func judgeDep(nums []byte) bool {

	temp := make(map[byte]bool)

	for i := range nums {
		if nums[i] == '.' {
			continue
		}

		if temp[nums[i]] {
			return true
		}

		temp[nums[i]] = true
	}

	return false
}

func isValidSudoku(board [][]byte) bool {

	if len(board) < 9 || len(board[0]) < 9 {
		return false
	}
	
	for i := range board {
		if judgeDep(board[i]) {
			return false
		}
	}

	for j := 0; j< len(board[0]); j ++ {
		temp := make([]byte,0)
		for i:=0; i < len(board); i ++ {
			temp = append(temp,board[i][j])
		}

		if judgeDep(temp) {
			return false
		}
	}

	for i := 0;i < 3; i ++ {
		for j := 0; j < 3; j++ {
			temp := make([]byte,0)
			startI,startJ := 3 * i, 3 * j

			for ti := startI;ti < startI + 3; ti ++ {
				for tj := startJ; tj < startJ + 3; tj ++ {
					temp = append(temp,board[ti][tj])
				}
			}

			if judgeDep(temp) {
				return false
			}
		} 
	}

	return true
}

func isValidSudokuIndex(board [][]byte, i, j int) bool {
	
	
	for k:= 0;k < 9 ; k++ {
		if board[i][k] == board[i][j] && k != j{
			return false
		}

		if board[k][j] == board[i][j] && k != i{
			return false
		}
	}

	startI,startJ := i/3 * 3, j/3 * 3

	for ti := startI;ti < startI + 3; ti ++ {
		for tj := startJ; tj < startJ + 3; tj ++ {
			if board[ti][tj] == board[i][j] && (ti != i || tj != j){
				return false
			}
		}
	}

	return true
}

var (
	chArray = []byte{'1','2','3','4','5','6','7','8','9'}
)

func solver(board [][]byte)bool{

	for i := 0; i < 9; i ++{
		for j := 0; j < 9; j ++ {
			if board[i][j] == '.' {
				for _, k := range chArray {
					board[i][j] = k
					
					if ok := isValidSudokuIndex(board,i,j); !ok{
						continue
					}

					if ok := solver(board); ok {
						return true
					}
				}
				board[i][j] = '.'
				return false
			}
		}
	}

	return true
}

func solveSudoku(board [][]byte)  {
	if len(board) < 9 || len(board[0]) < 9 {
		return
	}

	solver(board)
	
}
// @lc code=end

