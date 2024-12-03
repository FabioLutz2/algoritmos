package bucketsort

import selectionsort "algorithms/internal/sort/selectionSort"

func Sort(arr[] int) {
	length := len(arr)
	
	if length == 0 {
		return
	}
	
	maxValue := max(arr)
	
	bucket := make([][]int, (maxValue / 10) + 1)

	for i := 0; i < length; i++ {
		bucketIndex := arr[i] / 10
		bucket[bucketIndex] = append(bucket[bucketIndex], arr[i])
	}

	buckets := len(bucket)
	for i := 0; i < buckets; i++ {
		selectionsort.Sort(bucket[i])
	}

	index := 0
	for i := 0; i < buckets; i++ {
		for j := 0; j < len(bucket[i]); j++ {
			arr[index] = bucket[i][j]
			index++
		}
	}
}

func max(arr[] int) int {
	maxValue := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	
	return maxValue
}