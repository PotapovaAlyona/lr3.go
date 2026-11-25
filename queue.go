package main

import "fmt"

type Queue struct {
	Name  string
	Front *Node
	Rear  *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(data string) {
	newNode := &Node{Key: data}
	if q.Rear == nil {
		q.Front = newNode
		q.Rear = newNode
		return
	}
	q.Rear.Next = newNode
	q.Rear = newNode
}

func (q *Queue) Dequeue() (string, error) {
	if q.Front == nil {
		return "", fmt.Errorf("извлечение из пустой очереди")
	}
	val := q.Front.Key
	q.Front = q.Front.Next
	if q.Front == nil {
		q.Rear = nil
	}
	return val, nil
}

func (q *Queue) Peek() (string, error) {
	if q.Front == nil {
		return "", fmt.Errorf("очередь пуста")
	}
	return q.Front.Key, nil
}

func (q *Queue) Print() {
	fmt.Printf("Очередь %s: ", q.Name)
	for curr := q.Front; curr != nil; curr = curr.Next {
		fmt.Printf("%s ", curr.Key)
	}
	fmt.Println()
}
