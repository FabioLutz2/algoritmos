package quicksort

func partition(arr[] int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	j := low

	for j < high {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}

		j++
	}

	arr[i + 1], arr[high] = arr[high], arr[i + 1]

	return i + 1
}

func Sort(arr[] int, low int, high int) {
	if low < high {
		pivot := partition(arr, low, high)

		Sort(arr, low, pivot - 1)
		Sort(arr, pivot + 1, high)
	}
}