package main

import (
	binarysearch "algorithms/internal/search/binarySearch"
	exponentialsearch "algorithms/internal/search/exponentialSearch"
	interpolationsearch "algorithms/internal/search/interpolationSearch"
	jumpsearch "algorithms/internal/search/jumpSearch"
	ternarysearch "algorithms/internal/search/ternarySearch"
	bucketsort "algorithms/internal/sort/bucketSort"
	mergesort "algorithms/internal/sort/mergeSort"
	quicksort "algorithms/internal/sort/quickSort"
	radixsort "algorithms/internal/sort/radixSort"
	selectionsort "algorithms/internal/sort/selectionSort"
	shellsort "algorithms/internal/sort/shellSort"
	"fmt"
	"math/rand"
)

func main() {
	var algorithmChoice, orderChoice, length int

	fmt.Print(
		"\t### Bem vindo ###\n",
		"\n",
		"Escolha o algoritmo que deseja rodar:\n",
		"\n",
		"- SEARCH -\n",
		"1 - Binary Search\n",
		"2 - Interpolation Search\n",
		"3 - Jump Search\n",
		"4 - Exponential Search\n",
		"5 - Ternary Search\n",
		"\n",
		"- SORT -\n",
		"6 - Shell Sort\n",
		"7 - Merge Sort\n",
		"8 - Selection Sort\n",
		"9 - Bucket Sort\n",
		"10 - Radix Sort\n",
		"11 - Quick Sort\n",
		"Escolha: ",
	)
	fmt.Scanln(&algorithmChoice)

	fmt.Print(
		"\n",
		"Escolha se será ordenada ou não",
		"\n",
		"1 - Ordenada\n",
		"2 - Desordenada\n",
		"Escolha: ",
	)
	fmt.Scanln(&orderChoice)

	fmt.Print(
		"\n",
		"Escolha o limite do array",
		"\n",
		"Escolha: ",
	)
	fmt.Scanln(&length)

	var arr []int
	switch orderChoice {
	case 1:
		arr = generateSortedArray(length)
	case 2:
		arr = generateRandomArray(length)
	default:
		fmt.Println("Escolha da lista inválido")
		return
	}

	var result, searchValue int

	if algorithmChoice >= 1 && algorithmChoice <= 5 {
		fmt.Println("Array:")
		fmt.Println(arr)

		fmt.Print(
			"Escolha o valor a ser buscado\n",
			"Escolha: ",
		)
		fmt.Scanln(&searchValue)

		switch algorithmChoice {
		case 1:
			result = binarysearch.Search(arr, 0, length-1, searchValue)
		case 2:
			result = interpolationsearch.Search(arr, 0, length-1, searchValue)
		case 3:
			result = jumpsearch.Search(arr, 0, length-1, searchValue)
		case 4:
			result = exponentialsearch.Search(arr, 0, length-1, searchValue)
		case 5:
			result = ternarysearch.Search(arr, 0, length-1, searchValue)
		default:
			fmt.Println("Escolha de algoritmo inválido")
			return
		}

		fmt.Println("O valor está na posição:", result)
		return
	} else if algorithmChoice >= 6 && algorithmChoice <= 11 {
		fmt.Println("Array desordenado")
		fmt.Println(arr)

		switch algorithmChoice {
		case 6:
			shellsort.Sort(arr)
		case 7:
			mergesort.Sort(arr, 0, length - 1)
		case 8:
			selectionsort.Sort(arr)
		case 9:
			bucketsort.Sort(arr)
		case 10:
			radixsort.Sort(arr)
		case 11:
			quicksort.Sort(arr, 0, length - 1)
		}

		fmt.Println("Array organizado:")
		fmt.Println(arr)
		return
	} else {
		fmt.Println("Escolha de algoritmo inválido")
		return
	}
}

func generateSortedArray(end int) []int {
	var arr []int
	for i := 0; i < end; i++ {
		arr = append(arr, i)
	}

	return arr
}

func generateRandomArray(end int) []int {
	return rand.Perm(end)
}
