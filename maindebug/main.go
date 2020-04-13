package main

import (
	"fmt"
	"strings"
)

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
		i := n % 10
		n = n / 10

		sum += i
		pro = pro * i
	}

	res := pro - sum

	return res

}

func isMatch(s string, p string) bool {
	if p == ".*" {
		return true
	}

	if len(s) == 0 && len(p) == 2 && p[1:2] == "*" {
		return true
	}

	if len(s) <= 0 {
		return false
	}

	if len(s) > 0 && len(p) <= 0 {
		return false
	}

	if len(p) == 1 {
		if p[0:1] == "." || s[0:1] == p[0:1] {
			return isMatch(s[1:], p[1:])
		} else {
			return false
		}
	}

	if len(p) > 1 {
		if p[0:2] == ".*" {
			return isMatch(s, p[2:]) || isMatch(s[1:], p[2:]) || isMatch(s[1:], p)
		} else if p[0:1] == "." {
			return isMatch(s[1:], p[1:])
		} else if p[1:2] == "*" {
			return isMatch(s[1:], p[2:]) || isMatch(s[1:], p) || isMatch(s, p[2:])
		}
	}
	return false
}

func strNumAdd(str1, str2 string) string {
	if len(str1) > len(str2) {
		str1, str2 = str2, str1
	}

	ans := ""
	carry := 0
	leng1, leng2 := len(str1), len(str2)
	for i := 1; i <= leng1; i++ {
		sum := int(str1[leng1-i]-'0') + int(str2[leng2-i]-'0') + carry
		ans = fmt.Sprintf("%d", sum%10) + ans
		carry = sum / 10
	}

	for i := leng1 + 1; i <= leng2; i++ {
		sum := int(str2[leng2-i]-'0') + carry
		ans = fmt.Sprintf("%d", sum%10) + ans
		carry = sum / 10
	}

	if carry != 0 {
		ans = fmt.Sprintf("%d", carry) + ans
	}

	return ans
}

func product(nums string, ch byte) string {
	if ch == '0' {
		return "0"
	}

	carry := 0
	num := int(ch - '0')

	res := ""
	for i := len(nums) - 1; i >= 0; i-- {
		numi := int(nums[i] - '0')

		prod := num*numi + carry

		res = fmt.Sprintf("%d", prod%10) + res
		carry = prod / 10

	}

	if carry != 0 {
		res = fmt.Sprintf("%d", carry) + res
	}

	return res
}

func multiply(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	if num2 == "0" {
		return "0"
	}

	productMap := make(map[byte]string)

	ans := ""

	leng2 := len(num2)
	for i := leng2 - 1; i >= 0; i-- {
		v := productMap[num2[i]]
		if v == "" {
			v = product(num1, num2[i])
			productMap[num2[i]] = v
		}

		temp := v
		for j := 0; j < leng2-1-i; j++ {
			temp += "0"
		}

		ans = strNumAdd(ans, temp)
	}

	return ans
}

func main() {
	str := "/a/b/c/"

	strs := strings.Split(str, "/")

	for i := range strs {
		fmt.Printf("i: %d, s :%s\n", i, strs[i])
	}
}

func sort(arr [3]int) {
	arr[0] = 5
	arr[1] = 6
}

func sort2(arr *[]int) {
	(*arr)[0] = 3
	(*arr)[1] = 4
}

func handleNums(nums, flag []int) {
	if len(nums) == 0 || len(flag) == 0 {
		return
	}

	flag[0] = nums[0]

	min := 0
	for i := range nums {
		flag[i] = nums[min]

		if i > min && nums[i] < nums[min] {
			min = i
		}
	}

	return
}

func find(nums []int, start int, min int) int {

	for i := start; i < len(nums); i++ {
		if nums[i] > min && nums[i] < nums[start-1] {
			return i
		}
	}

	return -1
}

func find132pattern(nums []int) bool {

	flag := make([]int, len(nums))

	for i := 1; i < len(nums)-1; i++ {
		if flag[i] >= nums[i] {
			continue
		}

		if index := find(nums, i+1, flag[i]); index >= 0 {
			return true
		}
	}

	return false
}
