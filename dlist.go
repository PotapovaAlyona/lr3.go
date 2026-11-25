package main

import "fmt"

type DNode struct {
	Key  string
	Next *DNode
	Prev *DNode
}

type DoubleList struct {
	Name string
	Head *DNode
	Tail *DNode
}

func NewDoubleList() *DoubleList {
	return &DoubleList{}
}

func (dl *DoubleList) AddToBeginning(key string) {
	newNode := &DNode{Key: key}
	if dl.Head == nil {
		dl.Head = newNode
		dl.Tail = newNode
		return
	}
	newNode.Next = dl.Head
	dl.Head.Prev = newNode
	dl.Head = newNode
}

func (dl *DoubleList) AddToEnd(key string) {
	newNode := &DNode{Key: key}
	if dl.Head == nil {
		dl.Head = newNode
		dl.Tail = newNode
		return
	}
	newNode.Prev = dl.Tail
	dl.Tail.Next = newNode
	dl.Tail = newNode
}

func (dl *DoubleList) DeleteHead() {
	if dl.Head == nil {
		return
	}
	dl.Head = dl.Head.Next
	if dl.Head != nil {
		dl.Head.Prev = nil
	} else {
		dl.Tail = nil
	}
}

func (dl *DoubleList) DeleteTail() {
	if dl.Tail == nil {
		return
	}
	dl.Tail = dl.Tail.Prev
	if dl.Tail != nil {
		dl.Tail.Next = nil
	} else {
		dl.Head = nil
	}
}

func (dl *DoubleList) DeleteByValue(value string) bool {
	curr := dl.Head
	for curr != nil && curr.Key != value {
		curr = curr.Next
	}
	if curr == nil {
		return false
	}
	if curr == dl.Head {
		dl.DeleteHead()
	} else if curr == dl.Tail {
		dl.DeleteTail()
	} else {
		curr.Prev.Next = curr.Next
		curr.Next.Prev = curr.Prev
	}
	return true
}

func (dl *DoubleList) Print() {
	fmt.Printf("Двусвязный список %s: NULL <-> ", dl.Name)
	for curr := dl.Head; curr != nil; curr = curr.Next {
		fmt.Printf("%s <-> ", curr.Key)
	}
	fmt.Println("NULL")
}
