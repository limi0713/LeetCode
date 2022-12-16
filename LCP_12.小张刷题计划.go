/*
为了提高自己的代码能力，小张制定了 LeetCode 刷题计划，他选中了 LeetCode 题库中的 n 道题，编号从 0 到 n-1，
并计划在 m 天内按照题目编号顺序刷完所有的题目（注意，小张不能用多天完成同一题）。

在小张刷题计划中，小张需要用 time[i] 的时间完成编号 i 的题目。
此外，小张还可以使用场外求助功能，通过询问他的好朋友小杨题目的解法，可以省去该题的做题时间。
为了防止“小张刷题计划”变成“小杨刷题计划”，小张每天最多使用一次求助。

我们定义 m 天中做题时间最多的一天耗时为 T（小杨完成的题目不计入做题总时间）。请你帮小张求出最小的 T是多少。

示例 1：

    输入：time = [1,2,3,3], m = 2

    输出：3

    解释：第一天小张完成前三题，其中第三题找小杨帮忙；第二天完成第四题，并且找小杨帮忙。这样做题时间最多的一天花费了 3 的时间，并且这个值是最小的。

示例 2：

    输入：time = [999,999,999], m = 4

    输出：0

	解释：在前三天中，小张每天求助小杨一次，这样他可以在三天内完成所有的题目并不花任何时间。
	
限制：
    1 <= time.length <= 10^5
    1 <= time[i] <= 10000
    1 <= m <= 1000
*/

func minTime(time []int, m int) int {
    
    left, right := 0, 0
    
    for i := range time {
        right += time[i]
    }
    
    // 二分搜索，找到最小划分
    for left < right {
        mid := (left+right) >> 1
        
        /* 
        假设划分最大时间是mid(扣除了分组中最大时间)
        根据mid对time进行划分，看需要划分几组，表示需要多少天
        */
        arrCount := splitArr(time,mid)
        
        if arrCount > m {//当划分需要的天数>m时，表示mid太小，移动left
            left = mid+1
        }else{ // 当划分需要的天<=m时，表示mid值满足要求，移动right，以检验更小的mid值
            right = mid
        }
    }
    
    return left
}

func splitArr(time []int, limit int) int {
    count := 1 // 分组数
    
    /*
    curSum:当前分组和
    curMax:当前分组最大时间
    */
    curSum,curMax := 0,0
    
    for i := range time {
        if time[i] < curMax {
            /*
            当前时间<curMax时，判断time[i]能否被划分到当前分组
            */
            if curSum + time[i] - curMax <= limit {
                // 可以划分
                curSum += time[i]
            }else {
                //划分到下一组，当前分组结束
                count ++
                curSum,curMax = time[i],time[i]
            }
        } else {
             /*
            当前时间>=curMax时，time[i]会被扣除，判断time[i]能否被划分到当前分组
            */
            if curSum <= limit {
                //划分到当前分组，并更新当前分组curMax
                curSum += time[i]
                curMax = time[i]
            }else {
                count ++
                curSum,curMax = time[i],time[i]
            }
        }
    }
    
    return count
}