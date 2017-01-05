package stack

type Stack interface {
	Pop() interface{}
	Push(e interface{}) interface{}
	Size() int
}

type stack struct {
	elements []interface{}
}


func NewStack() Stack {
	return &stack{}
}

func (s *stack) Pop() interface{} {
	if len(s.elements) == 0 {
		return nil
	}

	value := s.elements[len(s.elements)-1]

	s.elements = s.elements[:len(s.elements)-1]

	return value
}

func (s *stack) Push(e interface{}) interface{} {
	s.elements = append(s.elements, e)

	return nil
}

func (s *stack) Size() int {
	return len(s.elements)
}
