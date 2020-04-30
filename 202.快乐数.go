/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 *
 * https://leetcode-cn.com/problems/happy-number/description/
 *
 * algorithms
 * Easy (59.98%)
 * Likes:    341
 * Dislikes: 0
 * Total Accepted:    74.7K
 * Total Submissions: 124.6K
 * Testcase Example:  '19'
 *
 * 编写一个算法来判断一个数 n 是不是快乐数。
 *
 * 「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到
 * 1。如果 可以变为  1，那么这个数就是快乐数。
 *
 * 如果 n 是快乐数就返回 True ；不是，则返回 False 。
 *
 *
 *
 * 示例：
 *
 * 输入：19
 * 输出：true
 * 解释：
 * 1^2 + 9^2 = 82
 * 8^2 + 2^2 = 68
 * 6^2 + 8^2 = 100
 * 1^2 + 0^2 + 0^2 = 1
 *
 *
 */

package main

import "fmt"

// @lc code=start
func isHappy(n int) bool {
	ans := make(map[int]bool)

	return findHappy(n, &ans)
}

func findHappy(n int, ans *map[int]bool) bool {
	if n == 1 {
		return true
	}

	if _, ok := (*ans)[n]; ok {
		return false
	}

	(*ans)[n] = false

	sum := 0
	for n > 0 {
		b := n % 10
		sum += b * b
		n /= 10
	}

	if sum == 1 {
		return true
	}

	return findHappy(sum, ans)
}

// @lc code=end

func main() {
	fmt.Println(isHappy(18))
}
