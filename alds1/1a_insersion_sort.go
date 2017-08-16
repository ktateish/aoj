package main

import "fmt"

func isort(a []int) {
	for i := 1; i < len(a); i++ {
		p(a)
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
	p(a)
}

func p(a []int) {
	sp := ""
	for i := 0; i < len(a); i++ {
		fmt.Printf("%s%d", sp, a[i])
		sp = " "
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	a := []int{}
	for i := 0; i < n; i++ {
		var v int
		fmt.Scanf("%d", &v)
		a = append(a, v)
	}

	isort(a)
}
