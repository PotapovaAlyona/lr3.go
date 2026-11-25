package main

import "fmt"

type Node struct {
	Key  string
	Next *Node
}

type ForwardList struct {
	Name string
	Head *Node
}

func NewForwardList() *ForwardList {
	return &ForwardList{}
}

func (fl *ForwardList) AddToBeginning(key string) {
	fl.Head = &Node{Key: key, Next: fl.Head}
}

func (fl *ForwardList) AddToEnd(key string) {
	newNode := &Node{Key: key}
	if fl.Head == nil {
		fl.Head = newNode
		return
	}
	curr := fl.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

func (fl *ForwardList) DeleteHead() {
	if fl.Head != nil {
		fl.Head = fl.Head.Next
	}
}

func (fl *ForwardList) DeleteTail() {
	if fl.Head == nil {
		return
	}
	if fl.Head.Next == nil {
		fl.Head = nil
		return
	}
	curr := fl.Head
	for curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next = nil
}

func (fl *ForwardList) DeleteByValue(value string) bool {
	if fl.Head == nil {
		return false
	}
	if fl.Head.Key == value {
		fl.DeleteHead()
		return true
	}
	curr := fl.Head
	for curr.Next != nil && curr.Next.Key != value {
		curr = curr.Next
	}
	if curr.Next != nil {
		curr.Next = curr.Next.Next
		return true
	}
	return false
}

func (fl *ForwardList) FindByValue(value string) *Node {
	for curr := fl.Head; curr != nil; curr = curr.Next {
		if curr.Key == value {
			return curr
		}
	}
	return nil
}

func (fl *ForwardList) Print() {
	fmt.Printf("Односвязный список %s: ", fl.Name)
	for curr := fl.Head; curr != nil; curr = curr.Next {
		fmt.Printf("%s -> ", curr.Key)
	}
	fmt.Println("NULL")
}
