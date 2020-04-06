/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 *
 * https://leetcode-cn.com/problems/wildcard-matching/description/
 *
 * algorithms
 * Hard (27.25%)
 * Likes:    304
 * Dislikes: 0
 * Total Accepted:    25K
 * Total Submissions: 91.8K
 * Testcase Example:  '"aa"\n"a"'
 *
 * 给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
 * 
 * '?' 可以匹配任何单个字符。
 * '*' 可以匹配任意字符串（包括空字符串）。
 * 
 * 
 * 两个字符串完全匹配才算匹配成功。
 * 
 * 说明:
 * 
 * 
 * s 可能为空，且只包含从 a-z 的小写字母。
 * p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
 * 
 * 
 * 示例 1:
 * 
 * 输入:
 * s = "aa"
 * p = "a"
 * 输出: false
 * 解释: "a" 无法匹配 "aa" 整个字符串。
 * 
 * 示例 2:
 * 
 * 输入:
 * s = "aa"
 * p = "*"
 * 输出: true
 * 解释: '*' 可以匹配任意字符串。
 * 
 * 
 * 示例 3:
 * 
 * 输入:
 * s = "cb"
 * p = "?a"
 * 输出: false
 * 解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
 * 
 * 
 * 示例 4:
 * 
 * 输入:
 * s = "adceb"
 * p = "*a*b"
 * 输出: true
 * 解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
 * 
 * 
 * 示例 5:
 * 
 * 输入:
 * s = "acdcb"
 * p = "a*c?b"
 * 输入: false
 * 
 */

// @lc code=start
func isMatch(s string, p string) bool {

	for i:=1;i<len(p);i++{
		if p[i] == '*' && p[i-1] == '*' {
			if i+1 < len(p) {
				p = p[:i] + p[i+1:]
				i --
			}
		}
	}

	if len(s) == 0 {
		return len(p) == 0 || p == "*"
	}

	if len(p) == 0 {
		return len(s) == 0
	}

	ans := make([][]bool,0)

	for i := 0; i < 2; i ++ {
		ans = append(ans,make([]bool,len(s)))
	}

	if p[0] == '*' || p[0] == '?' || p[0] == s[0] {
		ans[0][0] = true
	}

	for i := 1; i < len(s); i ++ {
		if p[0] == '*' {
			ans[0][i] = true
		}else {
			ans[0][i] = false
		}
	}

	for i := 1; i < len(p);i ++ {
		
		if p[i] == '*' {
			ans[i%2][0] = ans[(i-1)%2][0]
		}else if p[i] == s[0] || p[i] == '?' {
			ans[i%2][0] = isMatch("",p[:i])
		}else {
			ans[i%2][0] = false
		}

		for j := 1;j < len(s); j ++ {
			if s[j] == p[i] || p[i] == '?' {
				ans[i%2][j] = ans[(i-1)%2][j-1]
			}else if p[i] == '*' {
				ans[i%2][j] = ans[(i-1)%2][j] || ans[(i-1)%2][j-1] || ans[i%2][j-1]
			}else {
				ans[i%2][j] = false
			}
		}

	}

	return ans[(len(p)-1)%2][len(s)-1]


	// 递归迭代
	// resMap = make(map[string]bool)

	// resMap[s+"--"+p] =  solve(s,p)

	// return resMap[s+"--"+p]
}

var (
	resMap  map[string]bool
)

func solve(s, p string) bool {
	sp := s + "--" + p

	if v,ok := resMap[sp]; ok {
		return v
	}

	if len(s) == 0 {
		resMap[sp] =  len(p) == 0 || p == "*"
		return resMap[sp]
	}

	if len(p) == 0 {
		resMap[sp] = (len(s) == 0)
		return resMap[sp]
	}


	if p[0] == '?' {
		resMap[sp] = solve(s[1:],p[1:])
		return resMap[sp]
	} else if p[0] == '*' {
		resMap[sp] = solve(s,p[1:]) || solve(s[1:],p[1:]) || solve(s[1:],p)
		return resMap[sp]
	} else if p[0] == s[0] {

		index := 0
		for ;index < len(p);index ++ {
			if p[index] == '?' || p[index] == '*' {
				break
			}
		}
		
		if index >= len(p) {
			index = len(p)
		}

		if index <= len(s) && s[:index] == p[:index]{
			resMap[sp] = solve(s[index:], p[index:])
			return resMap[sp]
		}
	}

	resMap[sp] = false
	return resMap[sp]
}
// @lc code=end

