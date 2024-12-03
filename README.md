# Algoritmos de busca e de ordenação

## Binary Search

### Implementação

[Implementação do Binary Search](./internal/search/binarySearch/binary.go)

### Por que a lista deve estar ordenada

A lista deve estar ordenada porque ele compara a posição baseada no elemento do meio, considerando que os valores menores sempre estejam antes e valores maiores sempre estejam depois.

Caso a lista não esteja ordenada, não há como saber em qual metade o alvo estará.

## Interpolation Search

### Implementação

[Implementação do Interpolation Search](./internal/search/interpolationSearch/interpolation.go)

### Casos em que interpolation é mais eficiente que o binary search

O intepolation search segue uma implementação parecida com o binary search, mas que, ao invés de sempre ir para o meio do array, tenta fazer uma estimativa do valor imaginando como uma lista uniformemente distribuída ou próximo disso.

O interpolation é pensado para ser mais eficiente para casos uniformemente distribuídos, estimando uma posição mais próxima tentando ter menos comparações.

## Jump Search

### Implementação

[Implementação do Jump Search](./internal/search/jumpSearch/jump.go)

## Exponential Search

### Implementação

[Implementação do Exponential Search](./internal/search/exponentialSearch/exponential.go)

## Shell Sort

### Implementação

[Implementação do Shell Sort](./internal/sort/shellSort/shell.go)

## Merge Sort

### Implementação

[Implementação do Merge Sort](./internal/sort/mergeSort/merge.go)

### Dividir para conquistar

O merge sort usa o princípio de dividir para conquistar.

O dividir é a separação do array em subarrays, enquanto o conquistar é a junção final em um array de forma organizada.

## Selection Sort

### Implementação

[Implementação do Selection Sort](./internal/sort/selectionSort/selection.go)

## Bucket Sort

### Implementação

[Implementação do Bucket Sort](./internal/sort/bucketSort/bucket.go)

## Radix Sort

### Implementação

[Implementação do Radix Sort](./internal/sort/radixSort/radix.go)

## Quick Sort

### Implementação

[Implementação do Quick Sort](./internal/sort/quickSort/quick.go)

## Ternary Search

