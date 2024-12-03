package quicksort

import (
	"math/rand"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		array  []int
		expect []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{8, 3, 2, 9, 6, 5, 1, 0, 4, 7}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, test := range tests {
		Sort(test.array, 0, len(test.array)-1)
		for i := 0; i < len(test.array); i++ {
			if test.array[i] != test.expect[i] {
				t.Error(
					"\n",
					"Entrada:", test.array,
					"\n",
					"Esperado:", test.expect,
				)
			}
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	arr := rand.Perm(8_000)

	for i := 0; i < b.N; i++ {
		Sort(arr, 0, len(arr) - 1)
	}
}