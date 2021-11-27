package BucketSort

func BucketSort(arr []float64) {
	size := len(arr)
	bucket := make([]float64, 1)
	buckets := make([][]float64, size)
	// Create empty buckets
	for i := range arr {
		buckets[i] = bucket
	}
	// Add elements into buckets
	for i := range arr {
		bucketIndex := int(arr[i] * float64(size))
		buckets[bucketIndex] = append(buckets[bucketIndex], arr[i])
	}
	// Sort the elements of each bucket
	for i := range arr {
		InsertionSort(buckets[i])
	}
	// Get the sorted array
	index := 0
	for i := range arr {
		for j := range buckets[i] {
			arr[index] = buckets[i][j]
			index++
		}
	}
}

func InsertionSort(arr []float64) {
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
