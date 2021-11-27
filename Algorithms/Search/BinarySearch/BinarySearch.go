package BinarySearch

func BinarySearch1(arr []int, target, min, max int) int {
	for min <= max {
		mid := min + (max-min)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return -1
}

func BinarySearch2(arr []int, target, min, max int) int {
	if min <= max {
		mid := min + (max-min)/2
		if arr[mid] == target {
			return mid
		}
		if target < arr[mid] {
			return BinarySearch2(arr, target, min, mid-1)
		} else {
			return BinarySearch2(arr, target, mid+1, max)
		}
	}
	return -1
}
