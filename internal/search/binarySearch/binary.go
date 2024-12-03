package binarysearch

func Search(arr[] int, left int, right int, target int) int {
	for left <= right {
		var mid int = left + (right - left) / 2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}