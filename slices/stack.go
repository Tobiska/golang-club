package main

import "fmt"

type Stack struct {
	cnt []int
}

func NewStack() *Stack {
	return &Stack{
		cnt: make([]int, 0),
	}
}

func (s *Stack) Push(i int) {
	s.cnt = append(s.cnt, i)
}

func (s *Stack) Pop() (value int) {
	s.cnt, value = s.cnt[:len(s.cnt)-1], s.cnt[len(s.cnt)-1]
	return
}

func main() {
	st := NewStack()
	st.Push(1)
	st.Push(2)
	fmt.Println(st.Pop())
}
