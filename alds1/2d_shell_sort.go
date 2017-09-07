package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	defer stdout.Flush()
	n := readInt()
	a := readIntSlice(n)

	m, g, cnt := shellSort(a)
	println(m)
	printInts(g)
	println(cnt)
	for _, x := range a {
		println(x)
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
		printf("%s%v", sp, a[i])
		sp = " "
	}
	println()
}

var (
	readString func() string
	stdout     *bufio.Writer
)

func init() {
	readString = newReadString(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
}

func newReadString(ior io.Reader) func() string {
	r := bufio.NewScanner(ior)
	r.Split(bufio.ScanWords)

	return func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
}

func readInt() int {
	i, err := strconv.Atoi(readString())
	if err != nil {
		panic(err.Error())
	}
	return i
}

func readIntSlice(n int) []int {
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = readInt()
	}
	return b
}

func readLengthAndSlice() []int {
	n := readInt()
	return readIntSlice(n)
}

func printf(f string, args ...interface{}) (int, error) {
	return fmt.Fprintf(stdout, f, args...)
}

func println(args ...interface{}) (int, error) {
	return fmt.Fprintln(stdout, args...)
}
