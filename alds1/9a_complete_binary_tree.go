package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Heap struct {
	a []int
}

func NewHeap() *Heap {
	return &Heap{make([]int, 1)}
}

func parent(i int) int {
	return i / 2
}

func left(i int) int {
	return i * 2
}

func right(i int) int {
	return i*2 + 1
}

func (h *Heap) Push(k int) {
	h.a = append(h.a, k)
}

func (h *Heap) Print() {
	for i := 1; i < len(h.a); i++ {
		printf("node %d: key = %d, ", i, h.a[i])
		p := parent(i)
		l := left(i)
		r := right(i)
		if p > 0 {
			printf("parent key = %d, ", h.a[p])
		}
		if l < len(h.a) {
			printf("left key = %d, ", h.a[l])
		}
		if r < len(h.a) {
			printf("right key = %d, ", h.a[r])
		}
		println()
	}
}

func main() {
	defer stdout.Flush()
	n := readInt()
	h := NewHeap()
	for i := 0; i < n; i++ {
		k := readInt()
		h.Push(k)
	}
	h.Print()
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
