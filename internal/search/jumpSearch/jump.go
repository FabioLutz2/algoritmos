package jumpsearch

import "math"

func Search(arr[] int, start int, end int, target int) int {
	if start > end {
		return -1
	}

	var step int = int(math.Sqrt(float64(end - start + 1)))
	previous := start
	block := previous + step

	for previous <= end && arr[min(block, end)] < target {
		previous = block
		block += step
		if previous >= end {
			break
		}
	}

	for i := previous; i <= min(block, end); i++ {
		if arr[i] == target {
			return i
		}
	}

	return -1
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}