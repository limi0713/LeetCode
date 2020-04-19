/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 *
 * https://leetcode-cn.com/problems/interleaving-string/description/
 *
 * algorithms
 * Hard (39.61%)
 * Likes:    160
 * Dislikes: 0
 * Total Accepted:    12.5K
 * Total Submissions: 31.4K
 * Testcase Example:  '"aabcc"\n"dbbca"\n"aadbbcbcac"'
 *
 * 给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。
 * 
 * 示例 1:
 * 
 * 输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
 * 输出: false
 * 
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	lenS1, lenS2, lenS3 := len(s1), len(s2), len(s3)

	if lenS1 + lenS2 != lenS3 {
		return false
	}

	if lenS1 == 0 || lenS2 == 0 {
		return s2 == s3 || s1 == s3
	}

	indexS1, indexS2 :=0, 0
	
	for i := range s3 {

		if indexS1 < lenS1 && indexS2 < lenS2 && s1[indexS1] == s2[indexS2] && s1[indexS1] == s3[i] {

			s1Temp, s2Temp := "", ""
			if indexS1 + 1 < lenS1 {
				s1Temp = s1[indexS1+1:]

			}

			if indexS2 + 1 < lenS2{
				s2Temp = s2[indexS2+1:]
			}
			
			return isInterleave(s1Temp,s2[indexS2:],s3[i+1:]) || isInterleave(s1[indexS1:],s2Temp,s3[i+1:])
		}

		if indexS1 < lenS1 && s3[i] == s1[indexS1] {
			indexS1 ++
		}else if indexS2 < lenS2 && s3[i] == s2[indexS2] {
			indexS2 ++
		}else {
			return false
		}
	}

	return true
}
// @lc code=end

