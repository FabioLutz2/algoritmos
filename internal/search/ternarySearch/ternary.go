package ternarysearch

func Search(arr[] int, left int, right int, target int) int {
	if right >= left {
		mid1 := left + (right - left) / 3
		mid2 := right - (right - left) / 3

		if arr[mid1] == target {
			return mid1
		}

		if arr[mid2] == target {
			return mid2
		}

		if target < arr[mid1] {
			return Search(arr, left, mid1 - 1, target)
		} else if target > arr[mid2] {
			return Search(arr, mid2 + 1, right, target)
		} else {
			return Search(arr, mid1 + 1, mid2 - 1, target)
		}
	}

	return -1
}