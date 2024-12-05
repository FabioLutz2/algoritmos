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

### Listas desordenadas

Quando o interpolation é usado em listas desordenadas, não é possível de encontrar o elemento, pois o cálculo feito considera que seja uma lista ordenada, caso contrário, o cálculo não fará sentido.

### Casos em que interpolation é mais eficiente que o binary search

O intepolation search segue uma implementação parecida com o binary search, mas que, ao invés de sempre ir para o meio do array, tenta fazer uma estimativa do valor imaginando como uma lista uniformemente distribuída ou próximo disso.

O interpolation é pensado para ser mais eficiente para casos uniformemente distribuídos, estimando uma posição mais próxima tentando ter menos comparações.

## Jump Search

### Implementação

[Implementação do Jump Search](./internal/search/jumpSearch/jump.go)

### Tamanho ideal

O tamanho ideal para uma lista de tamanho n é √n (raiz de n). Ele cria um equilíbrio entre o tamanho dos saltos, para conseguir pular o maior número de valores, mas tendo o menor intervalo para demorar menos quando encontrar.

### Comparação com binary search

Cenário: lista ordenada e o elemento não está presente.

Jump Search:
* 100
    911ns, 1.109µs, 1.002µs
* 10.000
    3.591µs, 3.335µs, 3.408µs
* 1.000.000
    54.19µs, 82.422µs, 54.054µs

Binary Search:
* 100: 
    1.096µs, 665ns, 692ns
* 10.000:
    1.643µs, 1.127µs, 1.326µs
* 1.000.000:
    2.379µs, 2.165µs, 2.463µs

## Exponential Search

### Implementação

[Implementação do Exponential Search](./internal/search/exponentialSearch/exponential.go)

### Relação com jump search e binary search

O exponential search é uma combinação do jump search com o binary search.

A parte do jump search são os saltos feitos com valores exponenciais. A parte do binary search é a busca feita quando é encontrado o intervalo em que o elemento está presente.

### Desempenho

Cenário: lista ordenada e o elemento não está presente.

* 100:
    1.102µs, 878ns, 958ns
* 10.000:
    1.253µs, 1.143µs, 1.31µs
* 1.000.000:
    3.364µs, 3.245µs, 3.01µs

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

[Implenentação do Ternary Search](./internal/search/ternarySearch/ternary.go)

### Quando é mais eficiente que o binary search

O ternary search é mais eficiente que o binary search em situações em que a lista não é totalmente linear, como em caso de funções unimodais.

# Estabilidade

Algoritmos estáveis são os que preservam a ordem original de elementos iguais.

Estáveis: Radix Sort, Bucket Sort (se internamente também for), Merge Sort
Não estáveis: Quick Sort, Selection Sort, Shell Sort