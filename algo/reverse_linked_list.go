package main

func reverseNode(cur, prev *ListNode) *ListNode {
	if cur == nil {
		return prev
	}
	next := cur.Next
	cur.Next = prev
	return reverseNode(next, cur)
}

func reverseList(head *ListNode) *ListNode {
	return reverseNode(head, nil)
}

func reverseIter(head *ListNode) *ListNode {
	var cur, prev *ListNode
	cur = head
	for cur != nil {
		n := cur.Next
		cur.Next, prev, cur = prev, cur, n
	}
	return prev
}
