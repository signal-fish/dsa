package LinearSearch

import (
	"math/rand"
	"testing"
	"time"
)

func TestLinearSearch(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, random.Intn(100-10)+10)
	for i := range arr {
		arr[i] = random.Intn(100)
	}
	InsertionSort(arr)
	for _, value := range arr {
		result := LinearSearch(arr, value)
		if result == -1 {
			t.Fail()
		}
	}
}

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		indexItem := arr[i]
		j := i - 1
		for j >= 0 && indexItem < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = indexItem
	}
}
