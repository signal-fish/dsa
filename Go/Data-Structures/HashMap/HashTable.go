package HashTable

import (
	"fmt"
	"hash/fnv"
)

type Node struct {
	key   interface{}
	value interface{}
	next  *Node
}

type HashTable struct {
	data [256]*Node
}

func NewNode(key, value interface{}) *Node {
	return &Node{key, value, nil}
}

func NewHashTable() *HashTable {
	return &HashTable{}
}

func (table *HashTable) Add(key, value interface{}) {
	position := hash(key)
	if table.data[position] == nil {
		table.data[position] = NewNode(key, value)
		return
	}
	current := table.data[position]
	for current.next != nil {
		current = current.next
	}
	current.next = NewNode(key, value)
}

func (table *HashTable) Get(key interface{}) interface{} {
	position := hash(key)
	current := table.data[position]
	for current != nil {
		if current.key == key {
			return current.value
		}
		current = current.next
	}
	return -1
}

func (table *HashTable) Set(key, value interface{}) bool {
	position := hash(key)
	current := table.data[position]
	for current != nil {
		if current.key == key {
			current.value = value
			return true
		}
		current = current.next
	}
	return false
}

func (table *HashTable) Remove(key interface{}) bool {
	position := hash(key)
	if table.data[position] == nil {
		return false
	}
	if table.data[position].key == key {
		table.data[position] = table.data[position].next
		return true
	}
	current := table.data[position]
	for current.next != nil {
		if current.next.key == key {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func hash(key interface{}) uint8 {
	hash := fnv.New32a()
	hash.Write([]byte(fmt.Sprintf("%v", key)))
	return uint8(hash.Sum32() % 256)
}
