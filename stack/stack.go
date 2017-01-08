package stack

import "errors"

type StringStack interface {
	Pop() (string, error)
	Push(e string) interface{}
	Size() int
}

type stringStack struct {
	elements []string
}


func NewStack() StringStack {
	return &stringStack{}
}

func (s *stringStack) Pop() (string, error) {
	if len(s.elements) == 0 {
		return "", errors.New("No elements")
	}

	value := s.elements[len(s.elements)-1]

	s.elements = s.elements[:len(s.elements)-1]

	return value, nil
}

func (s *stringStack) Push(e string) interface{} {
	s.elements = append(s.elements, e)

	return nil
}

func (s *stringStack) Size() int {
	return len(s.elements)
}
