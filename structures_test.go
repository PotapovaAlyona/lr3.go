package main

import (
	"os"
	"strconv"
	"testing"
)

func TestDynamicArray(t *testing.T) {
	arr := NewDynamicArray(2)
	arr.Name = "m1"
	arr.AddToEnd("apple")
	arr.AddToEnd("banana")
	arr.AddToEnd("cherry") // вызов resizeIfNeeded

	if v, err := arr.GetElement(1); err != nil || v != "banana" {
		t.Fatalf("ожидалось banana, получено %s", v)
	}
	if err := arr.AddAtIndex(1, "orange"); err != nil {
		t.Fatal(err)
	}
	if err := arr.SetElement(1, "grape"); err != nil {
		t.Fatal(err)
	}
	if err := arr.RemoveAtIndex(1); err != nil {
		t.Fatal(err)
	}
	arr.Print()

	// Проверка граничных ошибок
	if _, err := arr.GetElement(-1); err == nil {
		t.Fatal("ожидалась ошибка выхода за границы")
	}
	if err := arr.AddAtIndex(99, "x"); err == nil {
		t.Fatal("ожидалась ошибка выхода за границы")
	}
	if err := arr.RemoveAtIndex(99); err == nil {
		t.Fatal("ожидалась ошибка выхода за границы")
	}
	if err := arr.SetElement(99, "x"); err == nil {
		t.Fatal("ожидалась ошибка выхода за границы")
	}
}

func TestListsStackQueue(t *testing.T) {
	// ForwardList
	fl := NewForwardList()
	fl.Name = "f1"
	fl.DeleteHead()
	fl.DeleteTail()
	if fl.DeleteByValue("none") {
		t.Fatal("не должно удалять из пустого")
	}
	fl.AddToEnd("b")
	fl.AddToBeginning("a")
	fl.AddToEnd("c")
	if fl.FindByValue("b") == nil || fl.FindByValue("z") != nil {
		t.Fatal("ошибка поиска в ForwardList")
	}
	fl.Print()
	if !fl.DeleteByValue("b") || !fl.DeleteByValue("a") {
		t.Fatal("ошибка удаления по значению")
	}
	fl.DeleteTail()
	fl.AddToEnd("only")
	fl.DeleteTail()

	// DoubleList
	dl := NewDoubleList()
	dl.Name = "d1"
	dl.DeleteHead()
	dl.DeleteTail()
	if dl.DeleteByValue("none") {
		t.Fatal("не должно удалять из пустого")
	}
	dl.AddToEnd("b")
	dl.AddToBeginning("a")
	dl.AddToEnd("c")
	dl.Print()
	if !dl.DeleteByValue("b") || !dl.DeleteByValue("a") || !dl.DeleteByValue("c") {
		t.Fatal("ошибка удаления в DoubleList")
	}
	dl.AddToEnd("only")
	dl.DeleteTail()

	// Stack
	st := NewStack()
	st.Name = "s1"
	if _, err := st.Pop(); err == nil {
		t.Fatal("ожидалась ошибка пустого стека")
	}
	if _, err := st.Peek(); err == nil {
		t.Fatal("ожидалась ошибка пустого стека")
	}
	st.Push("1")
	st.Push("2")
	if v, _ := st.Peek(); v != "2" {
		t.Fatal("ошибка Peek")
	}
	st.Print()
	if v, _ := st.Pop(); v != "2" {
		t.Fatal("ошибка Pop")
	}

	// Queue
	q := NewQueue()
	q.Name = "q1"
	if _, err := q.Dequeue(); err == nil {
		t.Fatal("ожидалась ошибка пустой очереди")
	}
	if _, err := q.Peek(); err == nil {
		t.Fatal("ожидалась ошибка пустой очереди")
	}
	q.Enqueue("1")
	q.Enqueue("2")
	if v, _ := q.Peek(); v != "1" {
		t.Fatal("ошибка Peek")
	}
	q.Print()
	if v, _ := q.Dequeue(); v != "1" {
		t.Fatal("ошибка Dequeue")
	}
}

func TestAVLTreeAndHashTables(t *testing.T) {
	// AVLTree (все повороты, поиск и удаление)
	tree := NewAVLTree()
	tree.Name = "t1"
	tree.Print()
	if tree.Remove(10) {
		t.Fatal("не должно удалять из пустого")
	}

	for _, k := range []int{30, 20, 10, 40, 50, 25, 35, 5, 15, 45, 55} {
		tree.Insert(k)
	}
	tree.Insert(30) // дубликат
	tree.Print()

	if !tree.Search(25) || tree.Search(999) {
		t.Fatal("ошибка поиска в АВЛ-дереве")
	}
	if !tree.Remove(5) || !tree.Remove(10) || !tree.Remove(30) || tree.Remove(999) {
		t.Fatal("ошибка удаления в АВЛ-дереве")
	}

	// ChainHashTable
	cht := NewChainHashTable(4)
	cht.Name = "cht1"
	cht.Insert("apple")
	cht.Insert("banana")
	cht.Insert("apple")
	if !cht.Contains("apple") || cht.Contains("missing") {
		t.Fatal("ошибка поиска в ChainHashTable")
	}
	cht.Print()
	if !cht.Remove("apple") || cht.Remove("apple") {
		t.Fatal("ошибка удаления в ChainHashTable")
	}

	// OpenHashTable
	oht := NewOpenHashTable(4)
	oht.Name = "oht1"
	oht.Insert("k1")
	oht.Insert("k2")
	oht.Insert("k3")
	if !oht.Contains("k2") || oht.Contains("missing") {
		t.Fatal("ошибка поиска в OpenHashTable")
	}
	oht.Print()
	if !oht.Remove("k2") || oht.Remove("k2") {
		t.Fatal("ошибка удаления в OpenHashTable")
	}
}

func TestDatabaseCLI(t *testing.T) {
	db := NewDatabase()
	db.ProcessCommand("MPUSHTOEND arr1 v1")
	db.ProcessCommand("FPUSH fl1 v1")
	db.ProcessCommand("LPUSH dl1 v1")
	db.ProcessCommand("SPUSH st1 v1")
	db.ProcessCommand("QPUSH q1 v1")
	db.ProcessCommand("TINSERT tr1 50")
	db.ProcessCommand("TINSERT tr1 30")
	db.ProcessCommand("TDEL tr1 30")
	db.ProcessCommand("HINSERT ht1 k1")
	db.ProcessCommand("PRINT arr1")
	db.ProcessCommand("PRINT tr1")
	db.ProcessCommand("UNKNOWN cmd")

	tmpFile := "test_go_db.json"
	defer os.Remove(tmpFile)
	if err := db.SaveToFile(tmpFile); err != nil {
		t.Fatal(err)
	}
	db2 := NewDatabase()
	db2.LoadFromFile(tmpFile)
	db2.ProcessCommand("PRINT dl1")
}

// --- Бенчмарки Go ---

func BenchmarkForwardListAdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fl := NewForwardList()
		for i := 0; i < 1500; i++ {
			fl.AddToEnd("val_" + strconv.Itoa(i))
		}
	}
}

func BenchmarkDoubleListAdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dl := NewDoubleList()
		for i := 0; i < 1500; i++ {
			dl.AddToEnd("val_" + strconv.Itoa(i))
		}
	}
}

func BenchmarkStackPushPop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		st := NewStack()
		for i := 0; i < 3750; i++ {
			st.Push("val")
		}
		for i := 0; i < 3750; i++ {
			_, _ = st.Pop()
		}
	}
}

func BenchmarkQueueEnqueueDequeue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		q := NewQueue()
		for i := 0; i < 3750; i++ {
			q.Enqueue("val")
		}
		for i := 0; i < 3750; i++ {
			_, _ = q.Dequeue()
		}
	}
}

func BenchmarkAVLTreeInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tree := NewAVLTree()
		for i := 0; i < 1500; i++ {
			tree.Insert(i)
		}
	}
}

func BenchmarkDynamicArrayAdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := NewDynamicArray(10)
		for i := 0; i < 2500; i++ {
			arr.AddToEnd("val_" + strconv.Itoa(i))
		}
	}
}

func BenchmarkHashTableInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ht := NewChainHashTable(500)
		for i := 0; i < 1500; i++ {
			ht.Insert("key_" + strconv.Itoa(i))
		}
	}
}

func TestSerializationJSONAndBinary(t *testing.T) {
	arr := NewDynamicArray(2)
	arr.Name = "arr_ser"
	arr.AddToEnd("alpha")
	arr.AddToEnd("beta")

	// Текстовый формат (JSON)
	if err := SaveJSON(arr, "arr_go.json"); err != nil {
		t.Fatal(err)
	}
	defer os.Remove("arr_go.json")

	var arrJSON DynamicArray
	if err := LoadJSON(&arrJSON, "arr_go.json"); err != nil {
		t.Fatal(err)
	}
	if arrJSON.Size != 2 || arrJSON.Data[1] != "beta" {
		t.Fatalf("ошибка десериализации JSON: %+v", arrJSON)
	}

	// Бинарный формат (Gob)
	if err := SaveBinary(arr, "arr_go.bin"); err != nil {
		t.Fatal(err)
	}
	defer os.Remove("arr_go.bin")

	var arrBin DynamicArray
	if err := LoadBinary(&arrBin, "arr_go.bin"); err != nil {
		t.Fatal(err)
	}
	if arrBin.Size != 2 || arrBin.Data[0] != "alpha" {
		t.Fatalf("ошибка бинарной десериализации: %+v", arrBin)
	}
}
