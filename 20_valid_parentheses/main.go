package problem20

import "errors"

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return false
	}
	stack := &Stack{storage: make([]byte, 0, len(s))}
	for i := 0; i < len(s); i++ {
		val := s[i]
		top, error := stack.Top()
		if val == byte('(') || val == byte('[') || val == byte('{') || error != nil {
			stack.Push(val)
		} else {
			if top == byte('(') && val == byte(')') || top == byte('[') && val == byte(']') || top == byte('{') && val == byte('}') {
				stack.Pop()
			} else {
				stack.Push(val)
			}
		}
	}
	return stack.IsEmpty()
}

type Stack struct {
	storage []byte
}

func (s *Stack) Pop() (byte, error) {
	if s.IsEmpty() {
		return byte(0), errors.New("stack is empty")
	}
	val := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return val, nil
}

func (s *Stack) Push(val byte) {
	s.storage = append(s.storage, val)
}

func (s *Stack) Top() (byte, error) {
	if s.IsEmpty() {
		return byte(0), errors.New("stack is empty")
	}
	return s.storage[len(s.storage)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.storage) == 0
}
