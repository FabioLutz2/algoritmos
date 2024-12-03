package interpolationsearch

func Search(arr[] int, low int, high int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	
	for low <= high && target >= arr[low] && target <= arr[high] {
		if arr[low] == arr[high] {
			if arr[low] == target {
				return low
			}
			break
		}

		var mid int = low + ((target - arr[low]) * (high - low)) / (arr[high] - arr[low])

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			low = mid + 1			
		} else {
			high = mid - 1
		}
	}

	return -1
}