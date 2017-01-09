package stack

import (
	"errors"
	"sync"
)

type StringStack interface {
	Pop() (string, error)
	Push(e string) interface{}
	Size() int
}

type stringStack struct {
	elements  []string
	writeLock sync.Mutex
	readLock sync.Mutex
}


func NewStack() StringStack {
	return &stringStack{}
}

func (s *stringStack) Pop() (string, error) {
	defer s.writeLock.Unlock()
	s.writeLock.Lock()
	if len(s.elements) == 0 {
		return "", errors.New("No elements")
	}

	value := s.elements[len(s.elements)-1]

	s.elements = s.elements[:len(s.elements)-1]

	return value, nil
}

func (s *stringStack) Push(e string) interface{} {
	defer s.readLock.Unlock()
	s.readLock.Lock()
	s.elements = append(s.elements, e)

	return nil
}

func (s *stringStack) Size() int {
	return len(s.elements)
}
