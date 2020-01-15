/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode-cn.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (37.26%)
 * Likes:    276
 * Dislikes: 0
 * Total Accepted:    81.7K
 * Total Submissions: 219.1K
 * Testcase Example:  '4'
 *
 * 实现 int sqrt(int x) 函数。
 * 
 * 计算并返回 x 的平方根，其中 x 是非负整数。
 * 
 * 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
 * 
 * 示例 1:
 * 
 * 输入: 4
 * 输出: 2
 * 
 * 
 * 示例 2:
 * 
 * 输入: 8
 * 输出: 2
 * 说明: 8 的平方根是 2.82842..., 
 * 由于返回类型是整数，小数部分将被舍去。
 * 
 * 
 */

// @lc code=start
func mySqrt(x int) int {
	
	max := x/2+1
	min := 1

	for min<=max{
		if min * min < x {
			min++
		}else if min * min == x {
			return min
		}else if min * min >x{
			break
		}
	}

	return min-1
}
// @lc code=end

