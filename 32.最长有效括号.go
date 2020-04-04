/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 *
 * https://leetcode-cn.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (30.02%)
 * Likes:    580
 * Dislikes: 0
 * Total Accepted:    49.5K
 * Total Submissions: 165K
 * Testcase Example:  '"(()"'
 *
 * 给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
 * 
 * 示例 1:
 * 
 * 输入: "(()"
 * 输出: 2
 * 解释: 最长有效括号子串为 "()"
 * 
 * 
 * 示例 2:
 * 
 * 输入: ")()())"
 * 输出: 4
 * 解释: 最长有效括号子串为 "()()"
 * 
 * 
 */

// @lc code=start

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValidParenthesesScan(s string) int {

	ansLeft, ansRight := 0, 0

	countLeft,countRight := 0, 0

	for i := 0; i<len(s); i ++ {
		ch := s[i]

		if ch == '('{
			countLeft ++
		} else {
			countRight ++
		}

		if countRight == countLeft {
			ansLeft = max(ansLeft,countLeft + countRight)
		}

		if countRight > countLeft {
			countLeft,countRight = 0, 0
		}

	}

	countLeft,countRight = 0, 0

	for i := len(s)-1; i>= 0; i -- {
		ch := s[i]

		if ch == '('{
			countLeft ++
		} else {
			countRight ++
		}

		if countRight == countLeft {
			ansRight = max(ansRight,countLeft + countRight)
		}

		if countRight < countLeft {
			countLeft,countRight = 0, 0
		}

	}

	return max(ansLeft,ansRight)
}

func longestValidParentheses(s string) int {

	ans := make([]int,len(s))

	res := 0

	for i := 1; i < len(s); i ++ {
		ch := s[i]

		if ch == '('{
			ans[i] = 0
		} else {
			
			if s[i-1] == '(' {
				if i - 2 >= 0 {
					ans[i] = ans[i-2] + 2
				}else{
					ans[i] = 2
				}

			} else {
				if i - ans[i-1] - 1 >= 0 {
					if s[i - ans[i-1] - 1] == '(' {
						if i - ans[i-1] - 2 >= 0 {
							ans[i] = ans[i - ans[i-1] - 2] + 2 + ans[i-1]
						}else {
							ans[i] = 2 + ans[i-1]
						}
					}
				}
			}
		}

		res = max(res,ans[i])
	}

	return res
}
// @lc code=end

