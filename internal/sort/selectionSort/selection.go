package selectionsort

func Sort(arr[] int) {
	length := len(arr)
	for i := 0; i < length - 1; i++ {
		min := i

		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		temp := arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
}