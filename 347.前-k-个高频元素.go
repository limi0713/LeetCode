/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start

type nc struct{
	key int
	count int
}

type ncst []nc

func (p ncst) Len() int { 
	return len(p) 
}

func (a ncst) Swap(i, j int) { 
	a[i], a[j] = a[j], a[i] 
}

func (p ncst) Less(i, j int) bool {
    return p[i].count > p[j].count
}

func topKFrequent(nums []int, k int) []int {
	var res []int
	nm := make(map[int]int)
	for _,num := range nums {
		if _,ok:=nm[num]; ok {
			nm[num]++
		} else {
			nm[num] = 1
		}
	}
	
	var ncs ncst
	for k, m := range nm {
		item := nc{
			key:k,
			count:m,
		}
		ncs = append(ncs,item)
	}
	sort.Sort(ncst(ncs))
	for i:=0;i<k;i++{
		res = append(res,ncs[i].key)
	}
	return res
}
// @lc code=end

