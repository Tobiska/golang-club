package main

import "fmt"

type Temperature struct {
	Val     int
	MinDays int
	MinVal  int
}

type DStack struct {
	cnt []Temperature
}

func NewDStack(size int) *DStack {
	return &DStack{
		cnt: make([]Temperature, 0, size),
	}
}

func (s *DStack) Push(val int) {
	headVal, headMinDays, isEmpty := s.Peek()
	if !isEmpty {
		s.cnt = append(s.cnt, val)
		s.minDays = append(s.minDays, 0)
		return
	}

	if headVal > val {
		s.minDays = append(s.minDays, 1)
	} else {
		if headMinDays == 0 {
			s.minDays = append(s.minDays, 0)
		} else {
			s.minDays = append(s.minDays, headMinDays+1)
		}
	}
	s.cnt = append(s.cnt, val)
}

func (s *DStack) Peek() (int, int, bool) {
	if len(s.cnt) <= 0 || len(s.minDays) <= 0 {
		return 0, 0, false
	}

	return s.cnt[len(s.cnt)-1], s.minDays[len(s.minDays)-1], true
}

func (s *DStack) GetMinDays() []int {
	return s.minDays
}

func dailyTemperatures(temperatures []int) []int {
	stack := NewDStack(len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		stack.Push(temperatures[i])
	}

	return stack.GetMinDays()
}

func main() {
	fmt.Println(dailyTemperatures([]int{1, 2, 3, 1, 1, 2}))
}
