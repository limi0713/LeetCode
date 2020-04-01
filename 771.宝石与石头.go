/*
 * @lc app=leetcode.cn id=771 lang=golang
 *
 * [771] 宝石与石头
 *
 * https://leetcode-cn.com/problems/jewels-and-stones/description/
 *
 * algorithms
 * Easy (82.00%)
 * Likes:    511
 * Dislikes: 0
 * Total Accepted:    72.8K
 * Total Submissions: 88.8K
 * Testcase Example:  '"aA"\n"aAAbbbb"'
 *
 *  给定字符串J 代表石头中宝石的类型，和字符串 S代表你拥有的石头。 S 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。
 * 
 * J 中的字母不重复，J 和 S中的所有字符都是字母。字母区分大小写，因此"a"和"A"是不同类型的石头。
 * 
 * 示例 1:
 * 
 * 输入: J = "aA", S = "aAAbbbb"
 * 输出: 3
 * 
 * 
 * 示例 2:
 * 
 * 输入: J = "z", S = "ZZ"
 * 输出: 0
 * 
 * 
 * 注意:
 * 
 * 
 * S 和 J 最多含有50个字母。
 * J 中的字符不重复。
 * 
 * 
 */

// @lc code=start
func numJewelsInStones(J string, S string) int {
    mapJ := handleStr(J)
    mapS := handleStr(S)

    count := 0
    for k,v := range mapS {
        if _,ok := mapJ[k];ok{
            count += v
        }
    }
    return count
}

func handleStr(str string)map[string]int{
    res := make(map[string]int)
    for i := range str{
        if _,ok := res[str[i:i+1]];ok{
           res[str[i:i+1]] ++
        }else{
            res[str[i:i+1]] = 1
        }
    }
    return res
}
// @lc code=end

