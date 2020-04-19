/*
 * @lc app=leetcode.cn id=466 lang=golang
 *
 * [466] 统计重复个数
 *
 * https://leetcode-cn.com/problems/count-the-repetitions/description/
 *
 * algorithms
 * Hard (29.66%)
 * Likes:    61
 * Dislikes: 0
 * Total Accepted:    3.8K
 * Total Submissions: 10.9K
 * Testcase Example:  '"acb"\n4\n"ab"\n2'
 *
 * 由 n 个连接的字符串 s 组成字符串 S，记作 S = [s,n]。例如，["abc",3]=“abcabcabc”。
 * 
 * 如果我们可以从 s2 中删除某些字符使其变为 s1，则称字符串 s1 可以从字符串 s2 获得。例如，根据定义，"abc" 可以从 “abdbec”
 * 获得，但不能从 “acbbe” 获得。
 * 
 * 现在给你两个非空字符串 s1 和 s2（每个最多 100 个字符长）和两个整数 0 ≤ n1 ≤ 10^6 和 1 ≤ n2 ≤
 * 10^6。现在考虑字符串 S1 和 S2，其中 S1=[s1,n1] 、S2=[s2,n2] 。
 * 
 * 请你找出一个可以满足使[S2,M] 从 S1 获得的最大整数 M 。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：
 * s1 ="acb",n1 = 4
 * s2 ="ab",n2 = 2
 * 
 * 返回：
 * 2
 * 
 * 
 */

// @lc code=start
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	lenS1, lenS2 := len(s1), len(s2)

	if lenS1 <= 0 || lenS2 <= 0 || lenS1 * n1 < lenS2 * n2 {
		return 0
	}

	indexS1, indexS2 := 0, 0
	ans := 0
	flag := make(map[int]([2]int))

	for indexS1 < lenS1 * n1 {
		if indexS1%lenS1 == lenS1 - 1 {
			if v, ok := flag[indexS2%lenS2]; ok {
				cycLen := indexS1 - v[0]
				cycCount := (n1*lenS1-indexS1)/cycLen
				cycS2Count := (indexS2 - v[1])/lenS2

				ans += cycCount * cycS2Count

				indexS1 += cycLen * cycCount

			}else{
				flag[indexS2%lenS2] = [2]int{indexS1,indexS2}
			}
		}

		if s1[indexS1%lenS1] == s2[indexS2%lenS2] {
			if indexS2%lenS2 == lenS2 - 1 {
				ans ++
			}
			indexS2 ++
		}

		indexS1 ++
	}

	return ans/n2
}
// @lc code=end

