package InsertionSort

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
