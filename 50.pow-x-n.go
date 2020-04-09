/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode-cn.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (34.27%)
 * Likes:    295
 * Dislikes: 0
 * Total Accepted:    64.5K
 * Total Submissions: 187.9K
 * Testcase Example:  '2.00000\n10'
 *
 * 实现 pow(x, n) ，即计算 x 的 n 次幂函数。
 * 
 * 示例 1:
 * 
 * 输入: 2.00000, 10
 * 输出: 1024.00000
 * 
 * 
 * 示例 2:
 * 
 * 输入: 2.10000, 3
 * 输出: 9.26100
 * 
 * 
 * 示例 3:
 * 
 * 输入: 2.00000, -2
 * 输出: 0.25000
 * 解释: 2^-2 = 1/2^2 = 1/4 = 0.25
 * 
 * 说明:
 * 
 * 
 * -100.0 < x < 100.0
 * n 是 32 位有符号整数，其数值范围是 [−2^31, 2^31 − 1] 。
 * 
 * 
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	powValues := make(map[int]float64)

	powValues[0] = 1
	powValues[1] = x

	powValues[2] = x * x

	absN := n
	if n < 0 {
		absN = - n
	}

	powFunc(x, absN, powValues)

	if n >= 0 {
		return powValues[absN]
	}

	return 1/powValues[absN]
}

func powFunc(x float64, n int,powValues map[int]float64) float64 {
	if v, ok := powValues[n]; ok {
		return v
	}

	powHalf := powFunc(x, n/2, powValues)

	res := powHalf * powHalf

	if n %2 == 1 {
		res *= x
	}

	powValues[n] = res

	return res
}

// @lc code=end

