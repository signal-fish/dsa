package RadixSort

func RadixSort(arr []int) {
	size := len(arr)
	max := GetMax(arr, size)
	for place := 1; max/place > 0; place *= 10 {
		CountingSort(arr, size, place)
	}
}

func GetMax(arr []int, n int) int {
	max := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func CountingSort(arr []int, size, place int) {
	output := make([]int, size+1)
	// Find out the maximum element
	max := arr[0]
	for i := 1; i < size; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	// Initialize count array with all 0
	count := make([]int, max+1)
	// Calculate count of elements
	for i := 0; i < size; i++ {
		count[(arr[i]/place)%10]++
	}
	// Calculate cumulative count
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	// Place the elements in sorted order
	for i := size - 1; i >= 0; i-- {
		output[count[(arr[i]/place)%10]-1] = arr[i]
		count[(arr[i]/place)%10]--
	}
	copy(arr, output)
}
