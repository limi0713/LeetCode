/*
 * @lc app=leetcode.cn id=65 lang=golang
 *
 * [65] 有效数字
 *
 * https://leetcode-cn.com/problems/valid-number/description/
 *
 * algorithms
 * Hard (18.79%)
 * Likes:    106
 * Dislikes: 0
 * Total Accepted:    11.7K
 * Total Submissions: 62K
 * Testcase Example:  '"0"'
 *
 * 验证给定的字符串是否可以解释为十进制数字。
 * 
 * 例如:
 * 
 * "0" => true
 * " 0.1 " => true
 * "abc" => false
 * "1 a" => false
 * "2e10" => true
 * " -90e3   " => true
 * " 1e" => false
 * "e3" => false
 * " 6e-1" => true
 * " 99e2.5 " => false
 * "53.5e93" => true
 * " --6 " => false
 * "-+3" => false
 * "95a54e53" => false
 * 
 * 说明: 我们有意将问题陈述地比较模糊。在实现代码之前，你应当事先思考所有可能的情况。这里给出一份可能存在于有效十进制数字中的字符列表：
 * 
 * 
 * 数字 0-9
 * 指数 - "e"
 * 正/负号 - "+"/"-"
 * 小数点 - "."
 * 
 * 
 * 当然，在输入中，这些字符的上下文也很重要。
 * 
 * 更新于 2015-02-10:
 * C++函数的形式已经更新了。如果你仍然看见你的函数接收 const char * 类型的参数，请点击重载按钮重置你的代码。
 * 
 */

// @lc code=start
var (
	stateFlag = []bool{false,false,true,true,true,false,true,false,false}

	// transferShip[i] 表示状态i 接受 +/1，0-9，. ，e 四种字符的状态转移
	transferShip = [][]int {
		{1,2,8,-1},
		{-1,2,8,-1},
		{-1,2,3,5},
		{-1,4,-1,5},
		{-1,4,-1,5},
		{7,6,-1,-1},
		{-1,6,-1,-1},
		{-1,6,-1,-1},
		{-1,3,-1,-1},
	}
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)

	if len(s) <= 0 {
		return false
	}

	state := 0
	for i := range s {
		switch {
		case s[i] == '-' || s[i] == '+' :
			state = transferShip[state][0]
		case s[i] >= '0' && s[i] <= '9' :
			state = transferShip[state][1]
		case s[i] == '.':
			state = transferShip[state][2]
		case s[i] =='e':
			state = transferShip[state][3]
		default :
			return false
		}

		if state == -1 {
			return false
		}
	}

	return stateFlag[state]
}
// @lc code=end

