/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 *
 * https://leetcode-cn.com/problems/edit-distance/description/
 *
 * algorithms
 * Hard (56.75%)
 * Likes:    678
 * Dislikes: 0
 * Total Accepted:    40.8K
 * Total Submissions: 71K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
 * 
 * 你可以对一个单词进行如下三种操作：
 * 
 * 
 * 插入一个字符
 * 删除一个字符
 * 替换一个字符
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：word1 = "horse", word2 = "ros"
 * 输出：3
 * 解释：
 * horse -> rorse (将 'h' 替换为 'r')
 * rorse -> rose (删除 'r')
 * rose -> ros (删除 'e')
 * 
 * 
 * 示例 2：
 * 
 * 输入：word1 = "intention", word2 = "execution"
 * 输出：5
 * 解释：
 * intention -> inention (删除 't')
 * inention -> enention (将 'i' 替换为 'e')
 * enention -> exention (将 'n' 替换为 'x')
 * exention -> exection (将 'n' 替换为 'c')
 * exection -> execution (插入 'u')
 * 
 * 
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	leng1, leng2 := len(word1), len(word2)

	if leng2 == 0 || leng1 == 0 {
		return leng1 + leng2
	}

	ans := make([][]int, 0)
	for i := 0; i < 2; i++ {
		arr := make([]int, leng1)
		ans = append(ans, arr)
	}

	if word1[0] != word2[0] {
		ans[0][0] = 1
	}

	for i := 1; i < leng1; i ++ {
		if word1[i] == word2[0] {
			ans[0][i] = i
		} else {
			ans[0][i] = ans[0][i-1] + 1
		}
	}

	// for i := 1; i < leng2; i ++ {
	// 	if word2[i] == word1[0] {
	// 		ans[i][0] = i
	// 	} else {
	// 		ans[i][0] = ans[i-1][0] + 1
	// 	}
	// }

	for i := 1; i < leng2; i ++ {
		if word2[i] == word1[0] {
			ans[i%2][0] = i
		} else {
			ans[i%2][0] = ans[(i-1)%2][0] + 1
		}

		for j := 1; j < leng1; j++ {
			if word1[j] == word2[i] {
				ans[i%2][j] = ans[(i-1)%2][j-1]
			} else {
				ans[i%2][j] = min(ans[(i-1)%2][j-1],min(ans[(i-1)%2][j],ans[i%2][j-1])) + 1
			}
		}
	}

	return ans[(leng2-1)%2][leng1-1]
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

