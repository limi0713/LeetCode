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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	midList = make([]int, 0)
)

func midDP(root *TreeNode) {
	if root == nil {
		return
	}

	midDP(root.Left)
	midList = append(midList, root.Val)
	midDP(root.Right)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	midDP(root)
	l := 0
	r := len(midList) - 1

	for l <= r {

		if midList[l] != midList[r] {
			return false
		}

		l++
		r--
	}

	return true
}

func subtractProductAndSum(n int) int {

	sum := 0
	pro := 1

	for n > 0 {
		i := n%10
		n = n/10

		sum += i
		pro = pro * i
	}

	res := pro - sum

	return res

}

func main() {
	subtractProductAndSum(234)
}
