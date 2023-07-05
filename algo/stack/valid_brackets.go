package main

type Stack struct {
	array []rune
}

func NewStack(size int) *Stack {
	return &Stack{
		array: make([]rune, 0, size),
	}
}

func (s *Stack) Len() int {
	return len(s.array)
}

func (s *Stack) String() string {
	return string(s.array)
}

func (s *Stack) Peek() rune {
	if len(s.array) > 0 {
		return s.array[len(s.array)-1]
	}

	return 0
}

func (s *Stack) Pop() rune {
	if len(s.array) == 0 {
		return 0
	}
	answer := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return answer
}

func (s *Stack) Push(str rune) {
	s.array = append(s.array, str)
}

func generateParenthesis(n int) []string {
	brackets := make([]string, 0, n*2)
	stack := NewStack(n * 2)
	genParenthesis(stack, n, 0, 0, &brackets)
	return brackets
}

// // n 2 op 2 cl 0
// // (
//    ((
//    ()

func genParenthesis(stack *Stack, n, open, close int, brackets *[]string) {

	if close == n && open == n {
		*brackets = append(*brackets, stack.String())
		return
	}

	if open < n {
		stack.Push('(')
		genParenthesis(stack, n, open+1, close, brackets)
		stack.Pop()
	}

	if close < open {
		stack.Push(')')
		genParenthesis(stack, n, open, close+1, brackets)
		stack.Pop()
	}
	return
}
