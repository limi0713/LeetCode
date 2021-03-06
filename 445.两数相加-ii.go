/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 *
 * https://leetcode-cn.com/problems/add-two-numbers-ii/description/
 *
 * algorithms
 * Medium (56.84%)
 * Likes:    183
 * Dislikes: 0
 * Total Accepted:    31.9K
 * Total Submissions: 56.2K
 * Testcase Example:  '[7,2,4,3]\n[5,6,4]'
 *
 * 给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
 * 
 * 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
 * 
 * 
 * 
 * 进阶：
 * 
 * 如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 8 -> 0 -> 7
 * 
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getNumFromList(ls *ListNode) string {
    numStr := ""

    for ls != nil {
        numStr += strconv.Itoa(ls.Val)
        ls = ls.Next
    }

    return numStr
}

func handleNumStrToList(numStr string) *ListNode {
    res := &ListNode{
        Val : int(numStr[0]-'0'),
        Next : nil,
    }

    p := res 

    for i := 1; i < len(numStr); i ++ {
        node := &ListNode{
            Val:int(numStr[i]-'0'),
        }

        p.Next = node
        p = p.Next
    }

    return res
}

func handle2NumStrAdd(numStr1, numStr2 string) string {

    res := ""

    index1, index2, carry := len(numStr1)-1, len(numStr2)-1, 0

    for index1 >=0 && index2 >=0 {
        sum := int(numStr1[index1] - '0') + int(numStr2[index2] - '0') + carry

        res = strconv.Itoa(sum%10) + res

        carry = sum/10
        index1 --
        index2 --
    }

    for index1 >= 0 {
        sum := int(numStr1[index1] - '0') + carry

        res = strconv.Itoa(sum%10) + res

        carry = sum/10
        index1 --
    }

    for index2 >= 0 {
        sum := int(numStr2[index2] - '0') + carry

        res = strconv.Itoa(sum%10) + res

        carry = sum/10
        index2 --
    }

    if carry != 0 {
        res = strconv.Itoa(carry) + res
    }

    return res
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    numStr := handle2NumStrAdd(getNumFromList(l1),getNumFromList(l2))

    return handleNumStrToList(numStr)
}
// @lc code=end

