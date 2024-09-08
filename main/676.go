package main

type byteMapNode struct {
	childNode [26]*byteMapNode
	end       bool
}

type MagicDictionary struct {
	byteMap *byteMapNode
}

func Constructor676() MagicDictionary {
	return MagicDictionary{
		byteMap: &byteMapNode{},
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		cur := this.byteMap
		for _, c := range word {
			idx := c - 'a'
			if cur.childNode[idx] == nil {
				cur.childNode[idx] = &byteMapNode{}
			}
			cur = cur.childNode[idx]
		}
		cur.end = true
	}
}

func (this *MagicDictionary) dfs(node *byteMapNode, word string, hasChange bool) bool {
	if word == "" {
		return hasChange && node.end
	}

	idx := word[0] - 'a'
	if node.childNode[idx] != nil && this.dfs(node.childNode[idx], word[1:], hasChange) {
		return true
	}

	if !hasChange {
		for _, child := range node.childNode {
			if child == nil {
				continue
			}

			if this.dfs(child, word[1:], true) {
				return true
			}
		}
	}

	return false
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return this.dfs(this.byteMap, searchWord, false)
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */

// 3137
func minimumOperationsToMakeKPeriodic(word string, k int) int {
	maxTimes := 0
	substrCountMap := make(map[string]int)
	for i := 0; i <= len(word)-k; i++ {
		subStr := word[i : i+k-1]
		substrCountMap[subStr]++
		maxTimes = max(maxTimes, substrCountMap[subStr])
	}
	return len(word)/k - maxTimes
}

// 2908
func minimumSum(nums []int) int {
	leftMin := make([]int, len(nums))
	leftMin[0] = 50
	for i := 1; i < len(nums); i++ {
		leftMin[i] = min(nums[i-1], leftMin[i-1])
	}

	rightMin := make([]int, len(nums))
	rightMin[len(nums)-1] = 50
	for i := len(nums) - 2; i >= 0; i-- {
		rightMin[i] = min(rightMin[i+1], nums[i+1])
	}

	minSum := 150
	for i := 1; i < len(nums)-1; i++ {
		if leftMin[i] < nums[i] && nums[i] > rightMin[i] {
			minSum = min(minSum, leftMin[i]+nums[i]+rightMin[i])
		}
	}
	if minSum == 150 {
		return -1
	}
	return minSum
}

// 3133
func minEnd(n int, x int) int64 {
	// n -= 1
	// i,j := 0, 0
	// for n>>j > 0 {
	//     if x >> i & 1 == 0 {
	//         x |= ((n>>j) & 1) << i
	//         j++
	//     }
	//     i++
	// }
	// return int64(x)

	n -= 1
	ns := numBS(n)
	xs := numBS(x)
	for len(xs) < len(ns) {
		xs += "0"
	}

	i, j := 0, 0
	res := ""
	for j < len(ns) {
		if i >= len(xs) || xs[i] == '0' {
			res += string(ns[j])
			j++
		} else {
			res += string(xs[i])
		}
		i++
	}
	for i < len(xs) {
		res += string(xs[i])
		i++
	}

	r := 0
	for i, ch := range res {
		if ch == '1' {
			r += 1 << i
		}
	}
	return int64(r)
}

func numBS(num int) string {
	s := ""
	for num > 0 {
		if num&1 == 0 {
			s += "0"
		} else {
			s += "1"
		}
		num = num >> 1
	}
	return s
}
