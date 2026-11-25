package main

import "fmt"

type ChainNode struct {
	Key  string
	Next *ChainNode
}

type ChainHashTable struct {
	Name      string
	Table     []*ChainNode
	TableSize int
	Count     int
}

func NewChainHashTable(size int) *ChainHashTable {
	if size <= 0 {
		size = 10
	}
	return &ChainHashTable{
		Table:     make([]*ChainNode, size),
		TableSize: size,
	}
}

func (ht *ChainHashTable) hashFunction(key string) int {
	h := 5381
	for _, c := range key {
		h = ((h << 5) + h) + int(c)
	}
	res := h % ht.TableSize
	if res < 0 {
		res = -res
	}
	return res
}

func (ht *ChainHashTable) Insert(key string) {
	if ht.Contains(key) {
		return
	}
	idx := ht.hashFunction(key)
	ht.Table[idx] = &ChainNode{Key: key, Next: ht.Table[idx]}
	ht.Count++
}

func (ht *ChainHashTable) Contains(key string) bool {
	idx := ht.hashFunction(key)
	for c := ht.Table[idx]; c != nil; c = c.Next {
		if c.Key == key {
			return true
		}
	}
	return false
}

func (ht *ChainHashTable) Remove(key string) bool {
	idx := ht.hashFunction(key)
	curr := ht.Table[idx]
	var prev *ChainNode
	for curr != nil {
		if curr.Key == key {
			if prev == nil {
				ht.Table[idx] = curr.Next
			} else {
				prev.Next = curr.Next
			}
			ht.Count--
			return true
		}
		prev = curr
		curr = curr.Next
	}
	return false
}

func (ht *ChainHashTable) Print() {
	fmt.Printf("Хеш-таблица (цепочки) %s:\n", ht.Name)
	for i := 0; i < ht.TableSize; i++ {
		if ht.Table[i] != nil {
			fmt.Printf("  [%d]: ", i)
			for c := ht.Table[i]; c != nil; c = c.Next {
				fmt.Printf("%s -> ", c.Key)
			}
			fmt.Println("NULL")
		}
	}
}

type OpenHashTable struct {
	Name      string
	Table     []*string
	TableSize int
	Count     int
}

func NewOpenHashTable(size int) *OpenHashTable {
	if size <= 0 {
		size = 10
	}
	return &OpenHashTable{
		Table:     make([]*string, size),
		TableSize: size,
	}
}

func (ht *OpenHashTable) hashFunction(key string) int {
	h := 5381
	for _, c := range key {
		h = ((h << 5) + h) + int(c)
	}
	res := h % ht.TableSize
	if res < 0 {
		res = -res
	}
	return res
}

func (ht *OpenHashTable) Insert(key string) {
	if ht.Count >= ht.TableSize || ht.Contains(key) {
		return
	}
	idx := ht.hashFunction(key)
	for i := 0; i < ht.TableSize; i++ {
		probe := (idx + i) % ht.TableSize
		if ht.Table[probe] == nil {
			val := key
			ht.Table[probe] = &val
			ht.Count++
			return
		}
	}
}

func (ht *OpenHashTable) Contains(key string) bool {
	idx := ht.hashFunction(key)
	for i := 0; i < ht.TableSize; i++ {
		probe := (idx + i) % ht.TableSize
		if ht.Table[probe] != nil && *ht.Table[probe] == key {
			return true
		}
	}
	return false
}

func (ht *OpenHashTable) Remove(key string) bool {
	idx := ht.hashFunction(key)
	for i := 0; i < ht.TableSize; i++ {
		probe := (idx + i) % ht.TableSize
		if ht.Table[probe] != nil && *ht.Table[probe] == key {
			ht.Table[probe] = nil
			ht.Count--
			return true
		}
	}
	return false
}

func (ht *OpenHashTable) Print() {
	fmt.Printf("Хеш-таблица (открытая адресация) %s:\n", ht.Name)
	for i := 0; i < ht.TableSize; i++ {
		if ht.Table[i] != nil {
			fmt.Printf("  [%d]: %s\n", i, *ht.Table[i])
		}
	}
}
