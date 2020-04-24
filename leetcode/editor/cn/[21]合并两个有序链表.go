//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
//
// 示例： 
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
// 
// Related Topics 链表



//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
    	return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
    if l1.Val<l2.Val {
    	head = l1
    	l1=l1.Next
	}else {
		head =l2
		l2=l2.Next
	}
	var p1 *ListNode
    p1=head
	for l1!=nil||l2!=nil{
		if l1 == nil {
			p1.Next=l2
			p1=p1.Next
			l2=l2.Next
		}else if l2==nil{
			p1.Next=l1
			p1=p1.Next
			l1=l1.Next
		}else {
			if l1.Val<l2.Val {
				p1.Next = l1
				p1=p1.Next
				l1=l1.Next
			}else {
				p1.Next =l2
				p1=p1.Next
				l2=l2.Next
			}
		}
	}
	return head
}
//leetcode submit region end(Prohibit modification and deletion)
