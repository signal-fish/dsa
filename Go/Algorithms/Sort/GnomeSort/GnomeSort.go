package GnomeSort

func GnomeSort(arr []int) {
	i := 0
	for i < len(arr)-1 {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			if i != 0 {
				i -= 1
			}
		} else {
			i += 1
		}
	}
}
