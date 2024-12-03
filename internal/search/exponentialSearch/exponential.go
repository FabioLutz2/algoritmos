package exponentialsearch

import binarysearch "algorithms/internal/search/binarySearch"

func Search(arr[] int, start int, end int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	if arr[start] == target {
		return 0
	}
	
	i := 1
	for i < end && arr[i] <= target {
		i *= 2
	}

	return binarysearch.Search(arr, i/2, min(i, end-1), target)
}