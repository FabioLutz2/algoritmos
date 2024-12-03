package jumpsearch

import "math"

func Search(arr[] int, start int, end int, target int) int {
	if end == 0 {
		return -1
	}

	var step int = int(math.Sqrt(float64(end)))
	previous := start
	block := step

	for previous < end && arr[min(block, end)-1] < target {
		previous = block
		block += step
		if previous >= end {
			return -1
		}
	}

	for i := previous; i < min(block, end); i++ {
		if arr[i] == target {
			return i
		}
	}

	return -1
}