/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 *
 * https://leetcode-cn.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (46.45%)
 * Likes:    243
 * Dislikes: 0
 * Total Accepted:    37.1K
 * Total Submissions: 79.9K
 * Testcase Example:  '"25525511135"'
 *
 * 给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
 * 
 * 示例:
 * 
 * 输入: "25525511135"
 * 输出: ["255.255.11.135", "255.255.111.35"]
 * 
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	ans := make([]string, 0)

	dfs(&ans,s,[]string{},0)

	return ans
}

func dfs(ans *[]string, s string, cur []string, index int){
	if len(cur) == 4 {
		temp := make([]string, len(cur))
		copy(temp, cur)
		*ans = append(*ans, strings.Join(temp, "."))
		return
	}

	if len(cur) == 3 {
		num, _ := strconv.Atoi(s[index:])
		if num > 255 || index>= len(s) || (index < len(s)-1 && s[index]=='0'){
			return
		}else {
			temp := make([]string, len(cur))
			copy(temp, cur)
			temp = append(temp, s[index:])
			*ans = append(*ans, strings.Join(temp, "."))
		}

		return
	}

	for i := 1; i < 4 && index + i <= len(s); i ++ {
		num,_ := strconv.Atoi(s[index:index+i])
		if num > 255 || (i > 1 && s[index]=='0') {
			break
		}

		temp := make([]string, len(cur))
		copy(temp, cur)
		temp = append(temp,s[index:index+i])
		dfs(ans,s,temp,index+i)
	}
}
// @lc code=end

