package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		array  []int
		target int
		expect int
	}{
		{[]int{}, 1, -1},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 6, 6},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, -1},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, -10, -1},
	}

	for _, test := range tests {
		result := Search(test.array, 0, len(test.array)-1, test.target)
		if result != test.expect {
			t.Error(
				"\n",
				"Entrada:", test.array,
				"\n",
				"Resultado:", result,
				"\n",
				"Esperado:", test.expect,
			)
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	var arr []int
	for i := 0; i < 1_000_000; i++ {
		arr = append(arr, i)
	}

	for i := 0; i < b.N; i++ {
		Search(arr, 0, len(arr)-1, 10000)
	}
}