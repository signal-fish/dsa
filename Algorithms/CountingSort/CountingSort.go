package CountingSort

func CountingSort(arr []int) {
	size := len(arr)
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

	// Store the count of each element
	for i := 0; i < size; i++ {
		count[arr[i]]++
	}
	// Store the cumulative count of each array
	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}
	// Find the index of each element of the original array in the count array
	for i := size - 1; i >= 0; i-- {
		output[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}
	// Step 6: Copy output array
	copy(arr, output)
}
