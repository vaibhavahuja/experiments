package main

import "fmt"

type MyMap[K comparable, V any] struct {
	hashMap map[K]V
}

func NewMap[K comparable, V any]() MyMap[K, V] {
	return MyMap[K, V]{
		hashMap: make(map[K]V),
	}
}

func (m *MyMap[K, V]) Add(key K, value V) {
	m.hashMap[key] = value
}

func (m *MyMap[K, V]) printMap() {
	for k, v := range m.hashMap {
		fmt.Println(k, v)
	}
}

func main() {
	m := NewMap[int, int]()
	m.Add(1, 2)
	mp := NewMap[string, string]()
	mp.Add("hey", "asdas")
	m.printMap()
	mp.printMap()
}
