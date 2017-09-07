package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func parent(i int) int {
	return i / 2
}

func left(i int) int {
	return i * 2
}

func right(i int) int {
	return i*2 + 1
}

func maxHeapify(a []int, i int) {
	l := left(i)
	r := right(i)
	h := len(a) - 1
	largest := -1

	if l <= h && a[l] > a[i] {
		largest = l
	} else {
		largest = i
	}
	if r <= h && a[r] > a[largest] {
		largest = r
	}

	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest)
	}
}

func main() {
	defer stdout.Flush()
	n := readInt()
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = readInt()
	}
	for i := n / 2; i > 0; i-- {
		maxHeapify(a, i)
	}
	for i := 1; i <= n; i++ {
		printf(" %d", a[i])
	}
	println()
}

var (
	stdin  *bufio.Reader
	stdout *bufio.Writer
)

func init() {
	stdin = bufio.NewReader(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
}

func scanf(f string, args ...interface{}) (int, error) {
	n, err := fmt.Fscanf(stdin, f, args...)
	for {
		r, _, err := stdin.ReadRune()
		if err != nil || !unicode.IsSpace(r) {
			stdin.UnreadRune()
			break
		}
	}
	return n, err
}

func readInt() int {
	var n int
	scanf("%d", &n)
	return n
}

func readIntSlice(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d", &a[i])
	}
	return a
}

func readLengthAndSlice() []int {
	n := readInt()
	return readIntSlice(n)
}

func readString() string {
	var s string
	scanf("%s", &s)
	return s
}

func printf(f string, args ...interface{}) (int, error) {
	return fmt.Fprintf(stdout, f, args...)
}

func println(args ...interface{}) (int, error) {
	return fmt.Fprintln(stdout, args...)
}
