package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}
	nswap := selectionSort(sort.IntSlice(A))
	p(A)
	fmt.Println(nswap)
}

func p(A []int) {
	sp := ""
	for _, a := range A {
		fmt.Printf("%s%d", sp, a)
		sp = " "
	}
	fmt.Println()
}

func selectionSort(s sort.Interface) int {
	nswap := 0
	for i := 0; i < s.Len(); i++ {
		minj := i
		for j := i; j < s.Len(); j++ {
			if s.Less(j, minj) {
				minj = j
			}
		}
		if minj != i {
			s.Swap(minj, i)
			nswap++
		}
	}
	return nswap
}
