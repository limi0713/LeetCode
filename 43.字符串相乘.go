/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 *
 * https://leetcode-cn.com/problems/multiply-strings/description/
 *
 * algorithms
 * Medium (41.98%)
 * Likes:    306
 * Dislikes: 0
 * Total Accepted:    53.7K
 * Total Submissions: 127.8K
 * Testcase Example:  '"2"\n"3"'
 *
 * 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
 * 
 * 示例 1:
 * 
 * 输入: num1 = "2", num2 = "3"
 * 输出: "6"
 * 
 * 示例 2:
 * 
 * 输入: num1 = "123", num2 = "456"
 * 输出: "56088"
 * 
 * 说明：
 * 
 * 
 * num1 和 num2 的长度小于110。
 * num1 和 num2 只包含数字 0-9。
 * num1 和 num2 均不以零开头，除非是数字 0 本身。
 * 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
 * 
 * 
 */

// @lc code=start

func product(nums string,ch byte) string {
	if ch == '0' {
		return "0"
	}

	carry := 0
	num := int(ch - '0')

	res := ""
	for i := len(nums) - 1; i >= 0; i -- {
		numi := int(nums[i] - '0')

		prod := num * numi + carry

		res = fmt.Sprintf("%d",prod%10) + res
		carry = prod/10

	}

	if carry != 0 {
		res = fmt.Sprintf("%d",carry) + res
	}

	return res
}

func multiply(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	if num2 == "0" {
		return "0"
	}

	productMap := make(map[byte]string)

	ans := ""

	leng2 := len(num2)
	for i := leng2 - 1; i >= 0; i -- {
		v := productMap[num2[i]]
		if v == "" {
			v = product(num1,num2[i])
			productMap[num2[i]] = v
		}

		temp := v
		for j := 0; j < leng2 - 1 - i; j ++ {
			temp += "0"
		}

		ans = strNumAdd(ans, temp)
	}

	return ans
}

func strNumAdd(str1, str2 string) string {
	if len(str1) > len(str2) {
		str1, str2 = str2, str1
	}

	ans := ""
	carry := 0
	leng1, leng2 := len(str1), len(str2)
	for i := 1; i <= leng1; i ++ {
		sum := int(str1[leng1-i] - '0') + int(str2[leng2-i] - '0') + carry
		ans = fmt.Sprintf("%d",sum%10) + ans
		carry = sum/10
	}

	for i := leng1 + 1; i <= leng2; i++ {
		sum := int(str2[leng2-i] - '0') + carry
		ans = fmt.Sprintf("%d",sum%10) + ans
		carry = sum/10
	}

	if carry != 0 {
		ans = fmt.Sprintf("%d",carry) + ans
	}

	return ans
}
// @lc code=end

