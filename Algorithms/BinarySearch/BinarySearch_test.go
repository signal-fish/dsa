package BinarySearch

import (
	"math/rand"
	"testing"
	"time"
)

// Bubble Sort
func SortArray(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		flag := true
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}

func TestBinarySearch1(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, random.Intn(100-10)+10)
	for i := range arr {
		arr[i] = random.Intn(100)
	}
	SortArray(arr)
	for _, value := range arr {
		result := BinarySearch1(arr, value, 0, len(arr)-1)
		if result == -1 {
			t.Fail()
		}
	}
}

func TestBinarySearch2(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, random.Intn(100-10)+10)
	for i := range arr {
		arr[i] = random.Intn(100)
	}
	SortArray(arr)
	for _, value := range arr {
		result := BinarySearch2(arr, value, 0, len(arr)-1)
		if result == -1 {
			t.Fail()
		}
	}
}
