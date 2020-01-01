/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 *
 * https://leetcode-cn.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (42.14%)
 * Likes:    395
 * Dislikes: 0
 * Total Accepted:    104.6K
 * Total Submissions: 247.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
 * 
 * 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
 * 
 * 你可以假设除了整数 0 之外，这个整数不会以零开头。
 * 
 * 示例 1:
 * 
 * 输入: [1,2,3]
 * 输出: [1,2,4]
 * 解释: 输入数组表示数字 123。
 * 
 * 
 * 示例 2:
 * 
 * 输入: [4,3,2,1]
 * 输出: [4,3,2,2]
 * 解释: 输入数组表示数字 4321。
 * 
 * 
 */

// @lc code=start
func plusOne(digits []int) []int {
	
	k := 1
	for i := len(digits)-1;i>=0;i--{
		v := digits[i]+k
		digits[i] = v%10
		k = v/10
		if k == 0 {
			return digits
		}
	}

	if k == 0 {
		return digits
	}

	res :=make([]int, 0)
	res = append(res,k)
	for i := range digits{
		res = append(res,digits[i])
	}
	return res
}
// @lc code=end

