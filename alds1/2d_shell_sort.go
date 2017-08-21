package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	n, _ := strconv.Atoi(in.Text())
	a := make([]int, n)
	for i := 0; i < n; i++ {
		in.Scan()
		a[i], _ = strconv.Atoi(in.Text())
	}
	//r := bufio.NewReader(os.Stdin)
	//var n int
	//fmt.Fscan(r, &n)
	//a := make([]int, n)
	//for i := 0; i < n; i++ {
	//	fmt.Fscan(r, &a[i])
	//}

	m, g, cnt := shellSort(a)
	fmt.Println(m)
	printInts(g)
	fmt.Println(cnt)
	for _, x := range a {
		fmt.Println(x)
	}
}

func shellSort(a []int) (int, []int, int) {
	cnt := 0
	g := []int{}
	for i := 1; i <= len(a); i = 3*i + 1 {
		g = append(g, i)
	}
	reverse(g)
	m := len(g)

	for _, gg := range g {
		cnt += insertionSort(a, gg)
	}

	return m, g, cnt
}

func insertionSort(a []int, g int) int {
	cnt := 0
	for k := 0; k < g; k++ {
		for i := k + g; i < len(a); i += g {
			key := a[i]
			j := i - g
			for j >= 0 && a[j] > key {
				a[j+g] = a[j]
				cnt++
				j -= g
			}
			a[j+g] = key
		}
	}
	return cnt
}

func reverse(a []int) {
	i := 0
	j := len(a) - 1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

func printInts(a []int) {
	sp := ""
	for i := 0; i < len(a); i++ {
		fmt.Printf("%s%v", sp, a[i])
		sp = " "
	}
	fmt.Println()
}
