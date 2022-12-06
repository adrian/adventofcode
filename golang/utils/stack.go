package utils

type Stack struct {
	contents []string
}

func (s *Stack) Push(c string) {
	s.contents = append(s.contents, c)
}

func (s *Stack) Pop() string {
	lastChar := s.contents[len(s.contents)-1]
	s.contents = s.contents[0 : len(s.contents)-1]
	return lastChar
}

func (s *Stack) Peek() string {
	return s.contents[len(s.contents)-1]
}