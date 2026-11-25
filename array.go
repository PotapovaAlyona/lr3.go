package main

import (
	"fmt"
	"strings"
)

type DynamicArray struct {
	Name     string
	Data     []string
	Capacity int
	Size     int
}

func NewDynamicArray(initialCapacity int) *DynamicArray {
	if initialCapacity <= 0 {
		initialCapacity = 10
	}
	return &DynamicArray{
		Data:     make([]string, initialCapacity),
		Capacity: initialCapacity,
		Size:     0,
	}
}

func (da *DynamicArray) resizeIfNeeded() {
	if da.Size >= da.Capacity {
		newCap := da.Capacity * 2
		newData := make([]string, newCap)
		copy(newData, da.Data[:da.Size])
		da.Data = newData
		da.Capacity = newCap
	}
}

func (da *DynamicArray) AddToEnd(value string) {
	da.resizeIfNeeded()
	da.Data[da.Size] = value
	da.Size++
}

func (da *DynamicArray) AddAtIndex(index int, value string) error {
	if index < 0 || index > da.Size {
		return fmt.Errorf("индекс вне диапазона")
	}
	da.resizeIfNeeded()
	for i := da.Size; i > index; i-- {
		da.Data[i] = da.Data[i-1]
	}
	da.Data[index] = value
	da.Size++
	return nil
}

func (da *DynamicArray) RemoveAtIndex(index int) error {
	if index < 0 || index >= da.Size {
		return fmt.Errorf("индекс вне диапазона")
	}
	for i := index; i < da.Size-1; i++ {
		da.Data[i] = da.Data[i+1]
	}
	da.Size--
	return nil
}

func (da *DynamicArray) GetElement(index int) (string, error) {
	if index < 0 || index >= da.Size {
		return "", fmt.Errorf("индекс вне диапазона")
	}
	return da.Data[index], nil
}

func (da *DynamicArray) SetElement(index int, value string) error {
	if index < 0 || index >= da.Size {
		return fmt.Errorf("индекс вне диапазона")
	}
	da.Data[index] = value
	return nil
}

func (da *DynamicArray) Print() {
	fmt.Printf("Массив %s: [ %s ]\n", da.Name, strings.Join(da.Data[:da.Size], ", "))
}
