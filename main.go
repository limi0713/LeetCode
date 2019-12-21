package main

import "fmt"

var topIndex int

var charStack []rune

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

var (
	chMap = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
)

func isValid(s string) bool {
	topIndex = 0
	charStack = make([]rune, 0)

	sr := []rune(s)

	for i := range sr {
		if sr[i] == rune('{') || sr[i] == rune('(') || sr[i] == rune('[') {
			push(sr[i])
		} else {
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

func countAndSay(n int) string {
	var str string
	var res string

	res = "1"

	for i := 1; i < n; i++ {
		str = res
		res = ""

		for j := 0; j < len(str); {
			c, k := 0, j

			for k < len(str) {

				if str[k] == str[j] {
					c++
					k++
				} else {
					break
				}
			}
			res += fmt.Sprintf("%d%c", c, str[j])

			j = k
		}
	}

	return res
}

func main() {
	fmt.Printf("%v", countAndSay(4))
}
