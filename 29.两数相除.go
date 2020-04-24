/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 *
 * https://leetcode-cn.com/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (19.41%)
 * Likes:    307
 * Dislikes: 0
 * Total Accepted:    41K
 * Total Submissions: 210.9K
 * Testcase Example:  '10\n3'
 *
 * 给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
 * 
 * 返回被除数 dividend 除以除数 divisor 得到的商。
 * 
 * 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) =
 * -2
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: dividend = 10, divisor = 3
 * 输出: 3
 * 解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
 * 
 * 示例 2:
 * 
 * 输入: dividend = 7, divisor = -3
 * 输出: -2
 * 解释: 7/-3 = truncate(-2.33333..) = -2
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 被除数和除数均为 32 位有符号整数。
 * 除数不为 0。
 * 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。
 * 
 * 
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == 0 || divisor == 0 {
		return 0
	}

	flag := 1
	if ( dividend > 0 && divisor < 0 ) || ( dividend < 0 && divisor > 0 ) {
		flag = -1
	}

	if dividend < 0 {
		dividend = -dividend
	}

	if divisor < 0 {
		divisor = -divisor
	}

	if dividend < divisor {
		return 0
	}

	res := 0
	d1,d2 := dividend,divisor
	if d1 >= d2 {
		res = 1
	}

	for ; d1 >= d2 + d2; d2 += d2{
		res += res
	}

	res += divide(dividend-d2,divisor)

	res = flag * res

	if res > 2147483647 || res <  -2147483648 {
		return 2147483647
	}

	return res
}
// @lc code=end

