package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func countingSort(a, b []int, k int) {
	c := make([]int, k+1)
	n := len(a) - 1
	for i := 1; i <= n; i++ {
		c[a[i]]++
	}

	for i := 1; i <= k; i++ {
		c[i] = c[i] + c[i-1]
	}

	for i := n; i > 0; i-- {
		b[c[a[i]]] = a[i]
		c[a[i]]--
	}
}

func main() {
	defer stdout.Flush()

	a := readLengthAndSlice()
	k := -1
	for _, x := range a {
		if x > k {
			k = x
		}
	}
	a = append([]int{0}, a...)
	b := make([]int, len(a))
	countingSort(a, b, k)
	sp := ""
	for i := 1; i < len(a); i++ {
		printf("%s%d", sp, b[i])
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
	i, _ := strconv.Atoi(readString())
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
