package Array

import (
	"errors"
	"fmt"
)

type Array struct {
	data   []int
	length uint
}

func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity),
		length: 0,
	}
}

func (array *Array) Len() uint {
	return array.length
}

func (array *Array) isIndexOutOfRange(index uint) bool {
	if ok := index >= uint(cap(array.data)); ok {
		return true
	}
	return false
}

func (array *Array) Find(index uint) (int, error) {
	if array.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return array.data[index], nil
}

func (array *Array) Insert(index uint, value int) error {
	if array.Len() == uint(cap(array.data)) {
		return errors.New("full array")
	}
	if index != array.length && array.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}
	for i := array.length; i > index; i-- {
		array.data[i] = array.data[i-1]
	}
	array.data[index] = value
	array.length++
	return nil
}

func (array *Array) InsertToTail(value int) error {
	return array.Insert(array.Len(), value)
}

func (array *Array) Delete(index uint) (int, error) {
	if array.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	value := array.data[index]
	for i := index; i < array.Len()-1; i++ {
		array.data[i] = array.data[i+1]
	}
	array.length--
	return value, nil
}

func (array *Array) Print() {
	var format string
	for i := uint(0); i < array.Len(); i++ {
		format += fmt.Sprintf("|%+v", array.data[i])
	}
	fmt.Println(format)
}
