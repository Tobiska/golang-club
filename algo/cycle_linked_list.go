package main

import "fmt"

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func main() {
	cycleNode1 := &ListNode1{}
	cycleNode2 := &ListNode1{}

	cycleNode1.Next, cycleNode2.Next = cycleNode2, cycleNode1
	fmt.Println(detectCycle(cycleNode1))
}

// O(n) - память, время
func detectCycle(head *ListNode1) *ListNode1 {
	saved := make(map[*ListNode1]struct{}, 0)
	for head != nil {
		if _, ok := saved[head]; ok {
			return head
		}
		saved[head] = struct{}{}
		head = head.Next
	}

	return nil
}
