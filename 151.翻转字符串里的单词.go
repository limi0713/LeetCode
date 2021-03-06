/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 *
 * https://leetcode-cn.com/problems/reverse-words-in-a-string/description/
 *
 * algorithms
 * Medium (37.17%)
 * Likes:    157
 * Dislikes: 0
 * Total Accepted:    58.9K
 * Total Submissions: 142.6K
 * Testcase Example:  '"the sky is blue"'
 *
 * 给定一个字符串，逐个翻转字符串中的每个单词。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入: "the sky is blue"
 * 输出: "blue is sky the"
 * 
 * 
 * 示例 2：
 * 
 * 输入: "  hello world!  "
 * 输出: "world! hello"
 * 解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
 * 
 * 
 * 示例 3：
 * 
 * 输入: "a good   example"
 * 输出: "example good a"
 * 解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 * 
 * 
 * 
 * 
 * 说明：
 * 
 * 
 * 无空格字符构成一个单词。
 * 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
 * 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 * 
 * 
 * 
 * 
 * 进阶：
 * 
 * 请选用 C 语言的用户尝试使用 O(1) 额外空间复杂度的原地解法。
 * 
 */

// @lc code=start
func reverseWords(s string) string {

	ans := ""

	start := -1 

	for i := 0; i < len(s); i ++ {
		if s[i] == ' ' && start == -1 {
			continue
		}

		if s[i] != ' ' && start == -1 {
			start = i
			continue
		}

		if s[i] == ' ' {
			ans = s[start:i] + " " + ans
			start = -1
		}
	}

	if start != -1 {
		ans = s[start:] + " " + ans
	}

	if len(ans) <= 0 {
		return ""
	}

	return ans[:len(ans)-1]
}
// @lc code=end

