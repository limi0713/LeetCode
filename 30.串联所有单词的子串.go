/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 *
 * https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/description/
 *
 * algorithms
 * Hard (30.06%)
 * Likes:    242
 * Dislikes: 0
 * Total Accepted:    28.3K
 * Total Submissions: 94.1K
 * Testcase Example:  '"barfoothefoobarman"\n["foo","bar"]'
 *
 * 给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
 * 
 * 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：
 * ⁠ s = "barfoothefoobarman",
 * ⁠ words = ["foo","bar"]
 * 输出：[0,9]
 * 解释：
 * 从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
 * 输出的顺序不重要, [9,0] 也是有效答案。
 * 
 * 
 * 示例 2：
 * 
 * 输入：
 * ⁠ s = "wordgoodgoodgoodbestword",
 * ⁠ words = ["word","good","best","word"]
 * 输出：[]
 * 
 * 
 */

// @lc code=start
func handleWordsToMap(words []string)map[string]int{

	res := make(map[string]int)

	for i := range words{
		res[words[i]] += 1
	}

	return res
}

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	wordsMap := handleWordsToMap(words)

	ans := make([]int,0)

	for i := 0; i < wordLen; i++{
		left,right := i,i
		
		tempMap := make(map[string]int)
		count := 0
		for right+wordLen <= len(s) {
			subStr := s[right:right+wordLen]
			right += wordLen

			if _,ok := wordsMap[subStr];!ok {
				left = right
				tempMap = make(map[string]int)
				count = 0
				continue
			}
			
			tempMap[subStr] += 1
			count ++

			for wordsMap[subStr] < tempMap[subStr] {
				leftSubStr := s[left:left+wordLen]
				tempMap[leftSubStr] -= 1
				left = left + wordLen
				count--
			}

			if count == len(words){
				ans = append(ans,left)
			}
		}
	}

	return ans
}
// @lc code=end

