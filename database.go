package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Database struct {
	Arrays       map[string]*DynamicArray
	ForwardLists map[string]*ForwardList
	DoubleLists  map[string]*DoubleList
	Stacks       map[string]*Stack
	Queues       map[string]*Queue
	Trees        map[string]*AVLTree
	ChainTables  map[string]*ChainHashTable
	OpenTables   map[string]*OpenHashTable
}

func NewDatabase() *Database {
	return &Database{
		Arrays:       make(map[string]*DynamicArray),
		ForwardLists: make(map[string]*ForwardList),
		DoubleLists:  make(map[string]*DoubleList),
		Stacks:       make(map[string]*Stack),
		Queues:       make(map[string]*Queue),
		Trees:        make(map[string]*AVLTree),
		ChainTables:  make(map[string]*ChainHashTable),
		OpenTables:   make(map[string]*OpenHashTable),
	}
}

func (db *Database) ProcessCommand(command string) {
	parts := strings.Fields(command)
	if len(parts) < 2 {
		fmt.Println("Неверный формат команды")
		return
	}
	cmd, name := parts[0], parts[1]

	switch cmd {
	case "MPUSHTOEND":
		if len(parts) < 3 {
			return
		}
		arr, ok := db.Arrays[name]
		if !ok {
			arr = NewDynamicArray(10)
			arr.Name = name
			db.Arrays[name] = arr
			fmt.Println("Создан новый массив:", name)
		}
		arr.AddToEnd(parts[2])

	case "FPUSH":
		if len(parts) < 3 {
			return
		}
		fl, ok := db.ForwardLists[name]
		if !ok {
			fl = NewForwardList()
			fl.Name = name
			db.ForwardLists[name] = fl
		}
		fl.AddToEnd(parts[2])

	case "LPUSH":
		if len(parts) < 3 {
			return
		}
		dl, ok := db.DoubleLists[name]
		if !ok {
			dl = NewDoubleList()
			dl.Name = name
			db.DoubleLists[name] = dl
		}
		dl.AddToEnd(parts[2])

	case "SPUSH":
		if len(parts) < 3 {
			return
		}
		st, ok := db.Stacks[name]
		if !ok {
			st = NewStack()
			st.Name = name
			db.Stacks[name] = st
		}
		st.Push(parts[2])

	case "QPUSH":
		if len(parts) < 3 {
			return
		}
		q, ok := db.Queues[name]
		if !ok {
			q = NewQueue()
			q.Name = name
			db.Queues[name] = q
		}
		q.Enqueue(parts[2])

	case "TINSERT":
		if len(parts) < 3 {
			return
		}
		val, err := strconv.Atoi(parts[2])
		if err != nil {
			return
		}
		t, ok := db.Trees[name]
		if !ok {
			t = NewAVLTree()
			t.Name = name
			db.Trees[name] = t
		}
		t.Insert(val)

	case "TDEL":
		if len(parts) < 3 {
			return
		}
		val, err := strconv.Atoi(parts[2])
		if err == nil {
			if t, ok := db.Trees[name]; ok {
				t.Remove(val)
			}
		}

	case "HINSERT":
		if len(parts) < 3 {
			return
		}
		ht, ok := db.ChainTables[name]
		if !ok {
			ht = NewChainHashTable(10)
			ht.Name = name
			db.ChainTables[name] = ht
		}
		ht.Insert(parts[2])

	case "PRINT":
		if a, ok := db.Arrays[name]; ok {
			a.Print()
		} else if f, ok := db.ForwardLists[name]; ok {
			f.Print()
		} else if d, ok := db.DoubleLists[name]; ok {
			d.Print()
		} else if s, ok := db.Stacks[name]; ok {
			s.Print()
		} else if q, ok := db.Queues[name]; ok {
			q.Print()
		} else if t, ok := db.Trees[name]; ok {
			t.Print()
		} else if h, ok := db.ChainTables[name]; ok {
			h.Print()
		} else if o, ok := db.OpenTables[name]; ok {
			o.Print()
		}

	default:
		fmt.Println("Неизвестная команда:", cmd)
	}
}

func (db *Database) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (db *Database) LoadFromFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Файл не найден, будет создан новый:", filename)
		return
	}
	_ = json.Unmarshal(data, db)
	// Восстановление связей Prev для двусвязного списка после JSON-загрузки
	for _, dl := range db.DoubleLists {
		var prev *DNode
		for curr := dl.Head; curr != nil; curr = curr.Next {
			curr.Prev = prev
			prev = curr
			dl.Tail = curr
		}
	}
}
