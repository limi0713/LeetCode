/*
 * @lc app=leetcode.cn id=2407 lang=golang
 *
 * [2407] 最长递增子序列 II
 */

// @lc code=start
type segment struct {
	l   int
	r   int
	max int
}

type segments []segment

func (s segments) build(index, l, r int) {
	s[index].l, s[index].r = l, r
	if l == r {
		return
	}

	mid := (l + r) >> 1
	s.build(index<<1, l, mid)
	s.build(index<<1+1, mid+1, r)
}

func (s segments) flush(index, num, val int) {
	if s[index].l == s[index].r {
		s[index].max = val
		return
	}

	mid := (s[index].l + s[index].r) >> 1
	if num <= mid {
		s.flush(index<<1, num, val)
	} else {
		s.flush(index<<1+1, num, val)
	}
	s[index].max = max(s[index<<1].max, s[index<<1+1].max)
}

func (s segments) query(index, l, r int) int {
	if l <= s[index].l && s[index].r <= r {
		return s[index].max
	}

	mid := (s[index].l + s[index].r) >> 1
	if r <= mid {
		return s.query(index<<1, l, r)
	} else if mid < l {
		return s.query(index<<1+1, l, r)
	}
	return max(s.query(index<<1, l, mid), s.query(index<<1+1, mid+1, r))
}

func lengthOfLIS(nums []int, k int) int {
	mx := 0
	for _, x := range nums {
		mx = max(mx, x)
	}
	segList := make(segments, mx*4)
	segList.build(1, 1, mx)

	for _, num := range nums {
		if num == 1 {
			segList.flush(1, 1, 1)
		} else {
			segList.flush(1, num, 1+segList.query(1, max(num-k, 1), num-1))
		}
	}
	return segList[1].max
}

// @lc code=end

