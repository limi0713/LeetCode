/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 *
 * https://leetcode-cn.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (61.18%)
 * Likes:    314
 * Dislikes: 0
 * Total Accepted:    64K
 * Total Submissions: 104.6K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
 * 
 * 示例:
 * 
 * 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * 输出:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 * 
 * 说明：
 * 
 * 
 * 所有输入均为小写字母。
 * 不考虑答案输出的顺序。
 * 
 * 
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {

	wordMap := make(map[string][]string)

	handlerWordsToMap(strs, wordMap)

	ans := make([][]string, 0)

	for _, v := range wordMap {
		ans = append(ans, v)
	}

	return ans
}

func handlerWordsToMap(strs []string, wordMap map[string][]string){

	for i := range strs {
		key := handleOneWord(strs[i])

		if _, ok := wordMap[key]; ok {
			wordMap[key] = append(wordMap[key], strs[i])
		} else {
			wordMap[key] = []string{strs[i]}
		}
	}
}

func handleOneWord(str string) string {
	chCount := make([]int,26)

	for i := range str {
		chCount[int(str[i]-'a')] ++
	}

	res := ""
	for i := range chCount {
		if chCount[i] > 0 {
			res += string('a' + i) + string(chCount[i])
		}
	}

	return res
}
// @lc code=end

