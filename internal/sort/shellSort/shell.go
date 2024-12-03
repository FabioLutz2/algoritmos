package shellsort

func Sort(arr[] int) {
	length := len(arr)
	for gap := length / 2; gap > 0; gap /= 2 {

		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i

			for j >= gap && arr[j - gap] > temp {
				arr[j] = arr[j - gap]
				j -= gap
			}
			
			arr[j] = temp
		}
	}
}