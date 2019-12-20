/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start

var topIndex int

var charStack []rune

var (
	chMap = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
)

func push(str rune) {

	if len(charStack) > topIndex {
		charStack[topIndex] = str
	} else {
		charStack = append(charStack, str)
	}
	topIndex++
}

func pop() rune {
	topIndex--
	return charStack[topIndex]
}

func isValid(s string) bool {
	topIndex = 0
	charStack = make([]rune, 0)

	sr := []rune(s)

	for i := range sr {
		if sr[i] == rune('{') || sr[i] == rune('(') || sr[i] == rune('[') {
			push(sr[i])
		} else {
			if topIndex <= 0{
				return false
			}
			lastOne := pop()
			if chMap[lastOne] != sr[i] {
				return false
			}
		}
	}

	if topIndex != 0 {
		return false
	}

	return true
}

// @lc code=end
