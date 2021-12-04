package CocktailSort

func CocktailSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		flag := true
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}
		for j := len(arr) - 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}
