/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 *
 * https://leetcode-cn.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (31.79%)
 * Likes:    158
 * Dislikes: 0
 * Total Accepted:    59.2K
 * Total Submissions: 185.4K
 * Testcase Example:  '"Hello World"'
 *
 * 给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
 * 
 * 如果不存在最后一个单词，请返回 0 。
 * 
 * 说明：一个单词是指由字母组成，但不包含任何空格的字符串。
 * 
 * 示例:
 * 
 * 输入: "Hello World"
 * 输出: 5
 * 
 * 
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	ss := strings.Split(s," ")
	if len(ss)<=0{
		return 0
	}

	return len(ss[len(ss)-1])
}
// @lc code=end

