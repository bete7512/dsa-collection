
type Stack struct {
	p    []string
	size int
}

func (s *Stack) Peek() string {
	if s.size == 0 {
		return ""
	}

	return s.p[s.size-1]
}

func (s *Stack) Pop() string {
	if s.size == 0 {
		return ""
	}

	last := s.p[s.size-1]
	s.p = s.p[:s.size-1]
	s.size--

	return last
}

func (s *Stack) Push(value string) {
	s.p = append(s.p, value)
	s.size++
}

func isValid(s string) bool {
	if s == "" {
		return true
	}
	runes := []rune(s)
	stack := Stack{}
	stack.Push(string(runes[0]))
	for i, r := range runes {
		if i == 0 {
			continue
		}
		if stack.Peek() == "(" {
			if string(r) == ")" {
				stack.Pop()
			} else {
				stack.Push(string(r))

			}
		} else if stack.Peek() == "{" {
			if string(r) == "}" {
				stack.Pop()
			} else {
				stack.Push(string(r))

			}
		} else if stack.Peek() == "[" {
			if string(r) == "]" {
				stack.Pop()
			} else {
				stack.Push(string(r))

			}
		} else {
			stack.Push(string(r))
		}
	}
	return len(stack.p) == 0
}
