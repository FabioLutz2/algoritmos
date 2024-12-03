package mergesort

func merge(arr[] int, left int, mid int, right int) {
	n1 := mid - left + 1
	n2 := right - mid

	leftArr := make([]int, n1)
	rightArr := make([]int, n2)

	for i := 0; i < n1; i++ {
		leftArr[i] = arr[left + i]
	}

	for i := 0; i < n2; i++ {
		rightArr[i] = arr[mid + 1 + i]
	}

	i := 0
	j := 0
	k := left

	for i < n1 && j < n2 {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = leftArr[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

func Sort(arr[] int, left int, right int) {
	if left < right {
		mid := left + (right - left) / 2

		Sort(arr, left, mid)
		Sort(arr, mid + 1, right)

		merge(arr, left, mid, right)
	}
}