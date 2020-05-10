package main

import (
	"fmt"
)

type MyT struct{ a int }

func (t MyT) pri1() { fmt.Println(" 1 ") }

// func (t *MyT) pri1() { // 自动生成
// 	(*t).pri1()
// }
func (t *MyT) pri2() { fmt.Println(" 2 ") }

// func (t MyT) pri2() { //不会自动生成
// 	(&t).pri2()
// }

// type MTI interface {
// 	pri1()
// 	pri2()
// }

type MyT2 struct{ a int }

func (t MyT2) pri1() {
	fmt.Printf("%p, %v\n", &t, t)
}

func (t *MyT2) pri2() {
	fmt.Printf("%p, %v\n", t, t)
	// fmt.Println("hello")
}

func main() {
	a := make([]int, 0)
	b := new(int)

	fmt.Println(a)
	fmt.Println(b)

	c := 1
	d := &c
	*d++
	b = &c

	fmt.Println(c)
	fmt.Println(*d)
	fmt.Println(b)
}

func kthSmallest(mat [][]int, k int) int {
	left, right := 0, 0
	for i := range mat {
		left += mat[i][0] // 最小和

		right += mat[i][len(mat[i])-1] // 最大和
	}

	min := left

	for left < right {
		mid := left + (right-left)/2 // 统计比mid小的和有多少个

		leftCount := 1 // 比mid小的个数，left比mid小，所以初始为1

		dfsk(mat, &leftCount, mid, min, k, 0)

		if leftCount >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func dfsk(mat [][]int, count *int, mid, min, k, index int) {
	if index >= len(mat) || *count > k || min > mid {
		return
	}

	dfsk(mat, count, mid, min, k, index+1)

	for i := 1; i < len(mat[index]); i++ {
		if min+mat[index][i]-mat[index][0] <= mid {
			*count += 1
			dfsk(mat, count, mid, min+mat[index][i]-mat[index][0], k, index+1)
		} else {
			break
		}
	}
}

func splitArray(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}

	index := make([]int, 0, len(nums)+1)

	return splitWithIndex(nums, len(nums)-1, &index)
}

func splitWithIndex(nums []int, end int, index *[]int) int {
	if end == 0 {
		return 1
	}

	if end == 1 {
		if gcd(nums[0], nums[1]) > 1 {
			return 1
		}

		*index = append(*index, 1)
		return 2
	}

	if gcd(nums[0], nums[end]) > 1 {
		*index = (*index)[: 0 : len(nums)+1]
		return 1
	}

	indexBefore := splitWithIndex(nums, end-1, index)

	if len(*index) == 0 && gcd(nums[0], nums[end]) > 1 {
		*index = (*index)[: 0 : len(nums)+1]
		return 1
	}

	if len(*index) > 0 && gcd(nums[end], nums[(*index)[len(*index)-1]]) > 1 {
		return indexBefore
	}

	*index = append(*index, end)
	return indexBefore + 1
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func reversePairs(nums []int) int {
	ans := 0

	merge(&nums, 0, len(nums)-1, &ans)

	return ans

}

func merge(nums *[]int, start, end int, ans *int) {
	if start >= end {
		return
	}

	if start+1 == end {
		if (*nums)[start] > (*nums)[end] {
			*ans += 1
			swap(nums, start, end)
		}
		return
	}

	mid := start + (end-start)/2
	merge(nums, start, mid, ans)
	merge(nums, mid+1, end, ans)

	i, j := mid, end

	temp := make([]int, end-start+1)
	indexOfTemp := len(temp) - 1

	for i >= start && j > mid {
		if (*nums)[i] > (*nums)[j] {
			*ans += j - mid
			temp[indexOfTemp] = (*nums)[i]
			i--

		} else {
			temp[indexOfTemp] = (*nums)[j]
			j--
		}

		indexOfTemp--
	}

	for i >= start {
		temp[indexOfTemp] = (*nums)[i]
		i--
		indexOfTemp--
	}

	for j > mid {
		temp[indexOfTemp] = (*nums)[j]
		j--
		indexOfTemp--
		// *ans += mid - start + 1
	}

	for i := range temp {
		(*nums)[start+i] = temp[i]
	}
}

func swap(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[j]
}

func findMinFibonacciNumbers(k int) int {
	arr := handle(k)

	return dp(arr, k)
}

func handle(k int) []int {
	arr := make([]int, 0)
	arr = append(arr, 1, 1)

	for i := 2; arr[i-1] < k; i++ {
		arr = append(arr, arr[i-1]+arr[i-2])
	}

	return arr
}

func dp(arr []int, k int) int {

	ans := 0

	i := len(arr) - 1
	for k > 0 {
		if k >= arr[i] {
			ans++
			k -= arr[i]
		}
		i--
	}

	return ans
}

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

func handleStrToMap(s string) map[byte]int {
	numsMap := make(map[byte]int)

	for i := range s {
		numsMap[s[i]] += 1
	}

	return numsMap
}

func judgeSubValid(s, t map[byte]int) bool {

	for k, v := range t {
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
	ans[0] = len(s)

	subStrMap := handleStrToMap(s[:len(t)])
	minAnsIndex := 0
	indexStart := 0
	for i := len(t); i <= len(s) && indexStart <= len(s)-len(t); {
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

	return s[minAnsIndex : minAnsIndex+ans[minAnsIndex]]
}

func combine(n int, k int) [][]int {
	if n < k {
		return [][]int{}
	}

	ans := make([][]int, 0)

	cur := make([]int, 0, k)

	dfs(&ans, n, k, cur)

	return ans
}

func dfs(ans *[][]int, n, k int, cur []int) {
	fmt.Printf("cur addr : %p ,", cur)
	fmt.Println(cur)
	if len(cur) == k {
		// temp := make([]int, 0, len(cur))
		// temp = append(temp, cur...)
		*ans = append(*ans, cur)
		return
	}

	start := 1
	if len(cur) > 0 {
		start = cur[len(cur)-1] + 1
	}
	for ; start <= n; start++ {
		// temp := make([]int,0)
		// temp = append(temp, cur...)
		// temp = append(temp, start)
		// dfs(ans,n,k,temp)

		leng := len(cur)
		cur = append(cur, start)
		dfs(ans, n, k, cur)
		cur = cur[:leng]
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
