/*
 * @lc app=leetcode.cn id=927 lang=golang
 *
 * [927] 三等分
 *
 * https://leetcode-cn.com/problems/three-equal-parts/description/
 *
 * algorithms
 * Hard (30.66%)
 * Likes:    22
 * Dislikes: 0
 * Total Accepted:    1.1K
 * Total Submissions: 3.7K
 * Testcase Example:  '[1,0,1,0,1]'
 *
 * 给定一个由 0 和 1 组成的数组 A，将数组分成 3 个非空的部分，使得所有这些部分表示相同的二进制值。
 * 
 * 如果可以做到，请返回任何 [i, j]，其中 i+1 < j，这样一来：
 * 
 * 
 * A[0], A[1], ..., A[i] 组成第一部分；
 * A[i+1], A[i+2], ..., A[j-1] 作为第二部分；
 * A[j], A[j+1], ..., A[A.length - 1] 是第三部分。
 * 这三个部分所表示的二进制值相等。
 * 
 * 
 * 如果无法做到，就返回 [-1, -1]。
 * 
 * 注意，在考虑每个部分所表示的二进制时，应当将其看作一个整体。例如，[1,1,0] 表示十进制中的 6，而不会是 3。此外，前导零也是被允许的，所以
 * [0,1,1] 和 [1,1] 表示相同的值。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：[1,0,1,0,1]
 * 输出：[0,3]
 * 
 * 
 * 示例 2：
 * 
 * 输出：[1,1,0,1,1]
 * 输出：[-1,-1]
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 3 <= A.length <= 30000
 * A[i] == 0 或 A[i] == 1
 * 
 * 
 * 
 * 
 */

// @lc code=start
func threeEqualParts(A []int) []int {
    if len(A)<3{
        return []int{-1,-1}
    }

    count := count1(A)
    if count%3 != 0 {
        return []int{-1,-1}
    }

    if count == 0 {
        return []int{0,2}
    }

    i,j := find1(A,count)
    c0 := count0(A)

    if judge(A,i,i+c0) && judge(A,j,j+c0){
        if handle(A,i+c0,j+c0+1){
            return []int{i+c0,j+c0+1}
        }
    }

    return []int{-1,-1}
}

func judge(A []int,i,j int)bool{

    for k := i+1;k<=j&&k<len(A);k++{
        if A[k] == 1 {
            return false
        }
    }

    return true
}

func find1(A []int,count int)(int,int){
    res := make([]int,0)
    c := count/3
    ct := 0
    for k:=0;k<len(A);k++{
        if A[k]== 1 {
            ct ++
        }

        if ct == c {
            res = append(res,k)
            ct = 0

            if len(res) == 2{
                break
            }
        }
    }

    return res[0],res[1]
}

func count1(A []int)int{
    count := 0

    for i := range A{
        if A[i] == 1 {
            count ++
        }
    }

    return count
}

func count0(A []int)int{
    c := 0
    for i := len(A)-1;i>=0;i--{
        if A[i] == 1 {
            return c
        }
        c ++
    }
    return c
}

func handle(A []int,i,j int)bool{

    str1 := ""
    str2 := ""
    str3 := ""
    index := 0 
    for ;index<=i;index++{
        str1 += fmt.Sprintf("%d",A[index])
    }
    for index = i+1;index<j;index++{
        str2 += fmt.Sprintf("%d",A[index])
    }
    for index = j;index<len(A);index++{
        str3 += fmt.Sprintf("%d",A[index])
    }

    

    str1 = strings.TrimLeft(str1,"0")
    str2 = strings.TrimLeft(str2,"0")
    str3 = strings.TrimLeft(str3,"0")

    if str1 == str2 && str1 == str3 {
        return true
    }

    return false
}
// @lc code=end

