/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文字符串 Ⅱ
 *
 * https://leetcode-cn.com/problems/valid-palindrome-ii/description/
 *
 * algorithms
 * Easy (36.77%)
 * Likes:    202
 * Dislikes: 0
 * Total Accepted:    40.7K
 * Total Submissions: 103.7K
 * Testcase Example:  '"aba"'
 *
 * 给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
 * 
 * 示例 1:
 * 
 * 
 * 输入: "aba"
 * 输出: True
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: "abca"
 * 输出: True
 * 解释: 你可以删除c字符。
 * 
 * 
 * 注意:
 * 
 * 
 * 字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
 * 
 * 
 */

// @lc code=start
func validPalindrome(s string) bool {
	left,right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left ++
			right --
		} else {
			return check(s,left+1,right) || check(s,left,right-1)
		}
	}

	return true
}

func check(s string,left,right int) bool {
	for left < right {
		if s[left] == s[right] {
			left ++
			right --
		}else{
			return false
		}
	}
	return true
}
// @lc code=end

