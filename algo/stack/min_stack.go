package main

import "fmt"

type Node struct {
	val  int
	min  int
	prev *Node
}

type MinStack struct {
	head *Node
}

func Constructor() MinStack {
	return MinStack{
		head: &Node{
			min:  int(^uint(0) >> 1),
			val:  int(^uint(0) >> 1),
			prev: nil,
		},
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (this *MinStack) Push(val int) {
	minVal := min(val, this.head.min)
	this.head = &Node{
		min:  minVal,
		val:  val,
		prev: this.head,
	}
}

func (this *MinStack) Pop() {
	if this.head.prev != nil {
		this.head = this.head.prev
	}
}

func (this *MinStack) Top() int {
	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.min
}

func main() {
	stack := Constructor()
	stack.Push(2)
	stack.Push(5)
	fmt.Println(stack.GetMin()) // 2
	stack.Push(1)
	fmt.Println(stack.GetMin()) // 1
	stack.Pop()
	fmt.Println(stack.GetMin()) // 2
}
