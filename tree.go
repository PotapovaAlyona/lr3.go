package main

import "fmt"

type TNode struct {
	Key    int
	Height int
	Left   *TNode
	Right  *TNode
}

type AVLTree struct {
	Name string
	Root *TNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func getHeight(n *TNode) int {
	if n == nil {
		return 0
	}
	return n.Height
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getBalance(n *TNode) int {
	if n == nil {
		return 0
	}
	return getHeight(n.Left) - getHeight(n.Right)
}

func rightRotate(y *TNode) *TNode {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	y.Height = maxInt(getHeight(y.Left), getHeight(y.Right)) + 1
	x.Height = maxInt(getHeight(x.Left), getHeight(x.Right)) + 1
	return x
}

func leftRotate(x *TNode) *TNode {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	x.Height = maxInt(getHeight(x.Left), getHeight(x.Right)) + 1
	y.Height = maxInt(getHeight(y.Left), getHeight(y.Right)) + 1
	return y
}

func (t *AVLTree) Insert(key int) {
	t.Root = insertNode(t.Root, key)
}

func insertNode(node *TNode, key int) *TNode {
	if node == nil {
		return &TNode{Key: key, Height: 1}
	}
	if key < node.Key {
		node.Left = insertNode(node.Left, key)
	} else if key > node.Key {
		node.Right = insertNode(node.Right, key)
	} else {
		return node
	}

	node.Height = 1 + maxInt(getHeight(node.Left), getHeight(node.Right))
	balance := getBalance(node)

	// Left Left
	if balance > 1 && key < node.Left.Key {
		return rightRotate(node)
	}
	// Right Right
	if balance < -1 && key > node.Right.Key {
		return leftRotate(node)
	}
	// Left Right
	if balance > 1 && key > node.Left.Key {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}
	// Right Left
	if balance < -1 && key < node.Right.Key {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}
	return node
}

func minValueNode(node *TNode) *TNode {
	curr := node
	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

func (t *AVLTree) Remove(key int) bool {
	var removed bool
	t.Root = removeNode(t.Root, key, &removed)
	return removed
}

func removeNode(node *TNode, key int, removed *bool) *TNode {
	if node == nil {
		return nil
	}
	if key < node.Key {
		node.Left = removeNode(node.Left, key, removed)
	} else if key > node.Key {
		node.Right = removeNode(node.Right, key, removed)
	} else {
		*removed = true
		if node.Left == nil || node.Right == nil {
			var temp *TNode
			if node.Left != nil {
				temp = node.Left
			} else {
				temp = node.Right
			}
			if temp == nil {
				return nil
			}
			*node = *temp
		} else {
			temp := minValueNode(node.Right)
			node.Key = temp.Key
			var dummy bool
			node.Right = removeNode(node.Right, temp.Key, &dummy)
		}
	}

	node.Height = 1 + maxInt(getHeight(node.Left), getHeight(node.Right))
	balance := getBalance(node)

	if balance > 1 && getBalance(node.Left) >= 0 {
		return rightRotate(node)
	}
	if balance > 1 && getBalance(node.Left) < 0 {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}
	if balance < -1 && getBalance(node.Right) <= 0 {
		return leftRotate(node)
	}
	if balance < -1 && getBalance(node.Right) > 0 {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}
	return node
}

func (t *AVLTree) Search(key int) bool {
	curr := t.Root
	for curr != nil {
		if key == curr.Key {
			return true
		} else if key < curr.Key {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	return false
}

func (t *AVLTree) Print() {
	if t.Root == nil {
		fmt.Printf("Дерево %s пустое\n", t.Name)
		return
	}
	fmt.Printf("АВЛ-дерево %s:\n", t.Name)
	printHorizontal(t.Root, 0)
}

func printHorizontal(node *TNode, level int) {
	if node == nil {
		return
	}
	printHorizontal(node.Right, level+1)
	for i := 0; i < level; i++ {
		fmt.Print("    ")
	}
	fmt.Printf("%d (h:%d)\n", node.Key, node.Height)
	printHorizontal(node.Left, level+1)
}
