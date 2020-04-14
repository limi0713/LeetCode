/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start

func handleStrToMap(s string) map[byte]int {
	numsMap := make(map[byte]int)
	
	for i := range s {
		numsMap[s[i]] += 1
	}

	return numsMap
}

func judgeSubValid(s, t map[byte]int) bool {

	for k,v := range t {
		if s[k] < v {
			return false
		}
	}

	return true
}

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	numsMap := handleStrToMap(t)

	ans := make([]int, len(s))
	ans[0] = len(s) * 2

	subStrMap := handleStrToMap(s[:len(t)])
	minAnsIndex := 0
	indexStart := 0
	for i := len(t); i <= len(s) && indexStart <= len(s)-len(t); {
		// if _, ok := numsMap[s[indexStart]]; !ok {
		// 	ans[indexStart] = len(s) * 2
		// 	indexStart ++
		// }

		if judgeSubValid(subStrMap, numsMap) {
			ans[indexStart] = i - indexStart
			if ans[indexStart] == len(t) {
				return s[indexStart:i]
			}

			if ans[indexStart] < ans[minAnsIndex] {
				minAnsIndex = indexStart
			}

			subStrMap[s[indexStart]] -= 1
			indexStart++
		} else {
			if i >= len(s) {
				break
			}
			subStrMap[s[i]] += 1
			i++
		}
	}

	if ans[minAnsIndex] > len(s) {
		return ""
	}

	return s[minAnsIndex : minAnsIndex+ans[minAnsIndex]]
}
// @lc code=end

