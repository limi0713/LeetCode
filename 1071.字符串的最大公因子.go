/*
 * @lc app=leetcode.cn id=1071 lang=golang
 *
 * [1071] 字符串的最大公因子
 *
 * https://leetcode-cn.com/problems/greatest-common-divisor-of-strings/description/
 *
 * algorithms
 * Easy (48.48%)
 * Likes:    93
 * Dislikes: 0
 * Total Accepted:    16.5K
 * Total Submissions: 28.6K
 * Testcase Example:  '"ABCABC"\n"ABC"'
 *
 * 对于字符串 S 和 T，只有在 S = T + ... + T（T 与自身连接 1 次或多次）时，我们才认定 “T 能除尽 S”。
 * 
 * 返回最长字符串 X，要求满足 X 能除尽 str1 且 X 能除尽 str2。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：str1 = "ABCABC", str2 = "ABC"
 * 输出："ABC"
 * 
 * 
 * 示例 2：
 * 
 * 输入：str1 = "ABABAB", str2 = "ABAB"
 * 输出："AB"
 * 
 * 
 * 示例 3：
 * 
 * 输入：str1 = "LEET", str2 = "CODE"
 * 输出：""
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= str1.length <= 1000
 * 1 <= str2.length <= 1000
 * str1[i] 和 str2[i] 为大写英文字母
 * 
 * 
 */

// @lc code=start
func gcdOfStrings(str1 string, str2 string) string {

	res := findModArray(len(str1),len(str2))

	for i:=range res{
		temp := str1[:res[i]]
		if handleStr(str1,temp) && handleStr(str2,temp){
			return temp
		}
	}

	return ""

}

func handleStr(str,subStr string)bool{

	if len(str)%len(subStr)!=0{
		return false
	}

	for i:=0;i<len(str);i+=len(subStr){
		
		if str[i:i+len(subStr)]!=subStr{
			return false
		}
	}

	return true
}

func findModArray(a,b int)[]int{
	if a > b {
		a, b = b, a
	}

	res := make([]int,0)
	for i:=a;i>0;i--{
		if a%i == 0 && b%i == 0{
			res = append(res,i)
		}
	}

	return res
}

// @lc code=end

