package BucketSort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestBucketSort(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr1 := make([]float64, random.Intn(100-10)+10)
	for i := range arr1 {
		arr1[i] = random.Float64()
	}
	InsertionSort(arr1)

	arr2 := make(sort.Float64Slice, len(arr1))
	copy(arr2, arr1)
	arr2.Sort()

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			t.Fail()
		}
	}
}
