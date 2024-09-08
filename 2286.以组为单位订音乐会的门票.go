/*
 * @lc app=leetcode.cn id=2286 lang=golang
 *
 * [2286] 以组为单位订音乐会的门票
 */

// @lc code=start
type segment struct {
	l, r int
	min  int
	sum  int
}

type BookMyShow struct {
	n, m        int
	segmentList []segment
}

func Constructor(n int, m int) BookMyShow {
	b := BookMyShow{
		n:           n,
		m:           m,
		segmentList: make([]segment, n*4),
	}
	b.buildSegmentList(1, 0, n-1)
	return b
}

func (this *BookMyShow) buildSegmentList(idx int, l, r int) {
	this.segmentList[idx].l, this.segmentList[idx].r = l, r
	if l == r {
		return
	}

	mid := (l + r) >> 1
	this.buildSegmentList(idx<<1, l, mid)
	this.buildSegmentList(idx<<1+1, mid+1, r)
}

func (this *BookMyShow) queryRow(idx, maxRow int, k int) int {
	if this.segmentList[idx].min > (this.m - k) {
		return -1
	}

	if this.segmentList[idx].l == this.segmentList[idx].r {
		return this.segmentList[idx].l
	}

	if this.segmentList[idx<<1].min <= (this.m - k) {
		return this.queryRow(idx<<1, maxRow, k)
	}
	mid := (this.segmentList[idx].l + this.segmentList[idx].r) >> 1
	if mid < maxRow {
		return this.queryRow(idx<<1+1, maxRow, k)
	}

	return -1
}

func (this *BookMyShow) querySum(idx, l, r int) int {
	if l <= this.segmentList[idx].l && this.segmentList[idx].r <= r {
		return this.segmentList[idx].sum
	}

	mid := (this.segmentList[idx].l + this.segmentList[idx].r) >> 1
	if mid < l {
		return this.querySum(idx<<1+1, l, r)
	}
	if r <= mid {
		return this.querySum(idx<<1, l, r)
	}
	return this.querySum(idx<<1, l, mid) + this.querySum(idx<<1+1, mid+1, r)
}

func (this *BookMyShow) book(idx, row, k int) {
	if this.segmentList[idx].l == this.segmentList[idx].r {
		this.segmentList[idx].min += k
		this.segmentList[idx].sum += k
		return
	}

	mid := (this.segmentList[idx].l + this.segmentList[idx].r) >> 1
	if row <= mid {
		this.book(idx<<1, row, k)
	} else {
		this.book(idx<<1+1, row, k)
	}

	this.segmentList[idx].min = min(this.segmentList[idx<<1].min, this.segmentList[idx<<1+1].min)
	this.segmentList[idx].sum = this.segmentList[idx<<1].sum + this.segmentList[idx<<1+1].sum
}

func (this *BookMyShow) Gather(k int, maxRow int) []int {
	row := this.queryRow(1, maxRow, k)
	if row == -1 {
		return nil
	}

	sum := this.querySum(1, row, row)
	this.book(1, row, k)
	return []int{row, sum}
}

func (this *BookMyShow) Scatter(k int, maxRow int) bool {
	sum := this.querySum(1, 0, maxRow)
	if (sum + k) > (maxRow+1)*this.m {
		return false
	}

	startRow := this.queryRow(1, maxRow, 1)
	for ; ; startRow++ {
		sum := this.m - this.querySum(1, startRow, startRow)
		if k <= sum {
			this.book(1, startRow, k)
			return true
		}

		k -= sum
		this.book(1, startRow, sum)
	}
}

/**
 * Your BookMyShow object will be instantiated and called as such:
 * obj := Constructor(n, m);
 * param_1 := obj.Gather(k,maxRow);
 * param_2 := obj.Scatter(k,maxRow);
 */

/**
 * Your BookMyShow object will be instantiated and called as such:
 * obj := Constructor(n, m);
 * param_1 := obj.Gather(k,maxRow);
 * param_2 := obj.Scatter(k,maxRow);
 */
// @lc code=end

