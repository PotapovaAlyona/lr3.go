package main

import "fmt"

type Stack struct {
	Name string
	Top  *Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(data string) {
	s.Top = &Node{Key: data, Next: s.Top}
}

func (s *Stack) Pop() (string, error) {
	if s.Top == nil {
		return "", fmt.Errorf("извлечение из пустого стека")
	}
	val := s.Top.Key
	s.Top = s.Top.Next
	return val, nil
}

func (s *Stack) Peek() (string, error) {
	if s.Top == nil {
		return "", fmt.Errorf("стек пустой")
	}
	return s.Top.Key, nil
}

func (s *Stack) Print() {
	fmt.Printf("Стек %s (сверху вниз): ", s.Name)
	for curr := s.Top; curr != nil; curr = curr.Next {
		fmt.Printf("%s ", curr.Key)
	}
	fmt.Println()
}
