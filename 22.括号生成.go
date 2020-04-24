/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (73.76%)
 * Likes:    833
 * Dislikes: 0
 * Total Accepted:    92.5K
 * Total Submissions: 125.4K
 * Testcase Example:  '3'
 *
 * 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
 * 
 * 例如，给出 n = 3，生成结果为：
 * 
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 * 
 * 
 */

// @lc code=start
func hanlde(resMap map[string]string,curStr string,open,close,n int){
	if open == close && len(curStr) == 2 *n{
		resMap[curStr]=curStr
		return
	}

	if open <n {
		hanlde(resMap,curStr+"(",open+1,close,n)
	}

	if close<open{
		hanlde(resMap,curStr+")",open,close+1,n)
	}
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	resMap := make(map[string]string)

	hanlde(resMap,"",0,0,n)	

	ans := make([]string,0)

	for _,v := range resMap{
		ans = append(ans,v)
	}

	return ans

}
// @lc code=end

