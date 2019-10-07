/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	str := "0123456789"
	la := len(a)
	lb := len(b)
	ia := la - 1
	ib := lb - 1
	var res string
	jw := 0
	for ia>=0||ib>=0 {
		ai := 0
		bi := 0
		if ia>=0 {
			ai = int(a[ia]-'0')
		}
		if ib >= 0 {
			bi = int(b[ib]-'0')
		}
		sum := ai+bi+jw
		jw = sum/2
		res = string(str[sum%2]) + res
		ia--
		ib-- 
	}
	if jw != 0{
		res = string(str[jw])+res
	}
	return res
}
// @lc code=end

