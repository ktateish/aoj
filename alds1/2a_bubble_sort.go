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
	nswap := bubbleSort(sort.IntSlice(A))
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

func bubbleSort(s sort.Interface) int {
	nswap := 0
	flag := true
	for i := 0; flag; i++ {
		flag = false
		for j := s.Len() - 1; j >= i+1; j-- {
			if s.Less(j, j-1) {
				s.Swap(j, j-1)
				nswap++
				flag = true
			}
		}
	}
	return nswap
}
