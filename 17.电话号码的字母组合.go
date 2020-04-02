/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (53.01%)
 * Likes:    643
 * Dislikes: 0
 * Total Accepted:    96K
 * Total Submissions: 181.1K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 * 
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 * 
 * 
 * 
 * 示例:
 * 
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 * 
 * 
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 * 
 */

// @lc code=start
var (
    chSet = []string{"","","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"}
)

var (
    res []string
)

func handle(s string){
    num,_ := strconv.Atoi(s)

    if num < 0 || num >=len(chSet){
        return
    }

    str := chSet[num]

    if len(res) <= 0{
        for _, v := range str{
            res = append(res,string(v))
        }
        return
    }

    temp := make([]string,0)

    for i:= range res{
        for _, v := range str{
            temp = append(temp,res[i]+string(v))
        }
    }

    res = temp
}

func letterCombinations(digits string) []string {
    res = make([]string,0)

    for i:=0;i<len(digits);i++{
        handle(digits[i:i+1])
    }

    return res
}
// @lc code=end

