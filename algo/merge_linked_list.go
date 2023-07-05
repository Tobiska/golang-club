package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var head *ListNode
	if list1.Val < list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}

	tail := head

	for {
		isList1 := list1 == nil
		isList2 := list2 == nil
		if isList1 && isList2 {
			return head
		} else if isList1 {
			tail.Next = &ListNode{
				Val: list2.Val,
			}
			list2 = list2.Next
		} else if isList2 {
			tail.Next = &ListNode{
				Val: list1.Val,
			}
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			tail.Next = &ListNode{
				Val: list2.Val,
			}
			list2 = list2.Next
		} else {
			tail.Next = &ListNode{
				Val: list1.Val,
			}
			list1 = list1.Next
		}

		tail = tail.Next
	}
}

func listPrint(list *ListNode) {
	for list != nil {
		fmt.Printf("%d ", list.Val)
		list = list.Next
	}
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	listPrint(mergeTwoLists(list1, list2))
}
