package BubbleSort

import (
	"log"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	log.Println("random:", random)
	arr1 := make([]int, random.Intn(100-10)+10)
	for i := range arr1 {
		arr1[i] = random.Intn(100)
	}

	arr2 := make(sort.IntSlice, len(arr1))
	copy(arr2, arr1)
	arr2.Sort()

	BubbleSort(arr1)
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			t.Fail()
		}
	}
}
