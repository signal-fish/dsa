package MergeSort

func Merge(arr []int, leftIndex, divideIndex, rightIndex int) {
	var tempArray1, tempArray2 []int
	for i := leftIndex; i <= divideIndex; i++ {
		tempArray1 = append(tempArray1, arr[i])
	}
	for i := divideIndex + 1; i <= rightIndex; i++ {
		tempArray2 = append(tempArray2, arr[i])
	}
	arrayIndex := leftIndex
	tempArray1Index := 0
	tempArray2Index := 0
	for tempArray1Index != len(tempArray1) && tempArray2Index != len(tempArray2) {
		if tempArray1[tempArray1Index] <= tempArray2[tempArray2Index] {
			arr[arrayIndex] = tempArray1[tempArray1Index]
			tempArray1Index += 1
		} else {
			arr[arrayIndex] = tempArray2[tempArray2Index]
			tempArray2Index += 1
		}
		arrayIndex += 1
	}
	for tempArray1Index < len(tempArray1) {
		arr[arrayIndex] = tempArray1[tempArray1Index]
		tempArray1Index += 1
		arrayIndex += 1

	}
	for tempArray2Index < len(tempArray2) {
		arr[arrayIndex] = tempArray2[tempArray2Index]
		tempArray2Index += 1
		arrayIndex += 1
	}
}

func MergeSort(arr []int, leftIndex, rightIndex int) {
	if leftIndex >= rightIndex {
		return
	}
	divideIndex := int((leftIndex + rightIndex) / 2)
	MergeSort(arr, leftIndex, divideIndex)
	MergeSort(arr, divideIndex+1, rightIndex)
	Merge(arr, leftIndex, divideIndex, rightIndex)
}
