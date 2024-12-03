package radixsort

func count(arr[] int, exp int) {
	length := len(arr)

	output := make([]int, length)
	count := [10]int {0}

	for i := 0; i < length; i++ {
		count[(arr[i] / exp) % 10]++
	}
	
	for i := 1; i < 10; i++ {
		count[i] += count[i - 1]
	}

	for i := length - 1; i >= 0; i-- {
		output[count[(arr[i] / exp) % 10] - 1] = arr[i]
		count[(arr[i] / exp) % 10]--
	}

	for i := 0; i < length; i++ {
		arr[i] = output[i]
	}
}

func Sort(arr[] int) {
	length := len(arr)

	if length == 0 {
		return
	}

	maxNumber := max(arr, length)

	for exp := 1; maxNumber / exp > 0; exp *= 10 {
		count(arr, exp)
	}
}

func max(arr[] int, length int) int {
	maxValue := arr[0]

	for i := 1; i < length; i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	return maxValue
}