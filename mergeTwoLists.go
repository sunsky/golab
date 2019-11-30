package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	out := l1
	in := l2
	ret := out
	for out != nil {
		for in != nil {
			if out == nil {
				out = in
				in = nil
			} else if out.Val <= in.Val {
				out = out.Next
			} else {
				tmp := out.Next
				out.Next = &ListNode{out.Val, tmp}
				out.Val = in.Val
				in = in.Next
			}

		}
	}
	return ret
}
func main() {
	fmt.Printf("%#v \n", mergeTwoLists(&ListNode{1, &ListNode{2, &ListNode{3, nil}}}, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}))
}
