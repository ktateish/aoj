package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

func (h *Heap) Insert(k int) {
	var shiftUp func(int)
	shiftUp = func(i int) {
		p := parent(i)
		if p <= 0 || h.a[p] > h.a[i] {
			return
		}
		h.a[p], h.a[i] = h.a[i], h.a[p]
		shiftUp(p)
	}
	h.a = append(h.a, k)
	shiftUp(len(h.a) - 1)
}

func (h *Heap) ExtractMax() int {
	var maxHeapify func(int)
	maxHeapify = func(i int) {
		l := left(i)
		r := right(i)
		n := len(h.a) - 1
		largest := -1

		if l <= n && h.a[l] > h.a[i] {
			largest = l
		} else {
			largest = i
		}
		if r <= n && h.a[r] > h.a[largest] {
			largest = r
		}

		if largest != i {
			h.a[i], h.a[largest] = h.a[largest], h.a[i]
			maxHeapify(largest)
		}
	}
	res := h.a[1]
	n := len(h.a) - 1
	h.a[1] = h.a[n]
	h.a = h.a[:n]
	maxHeapify(1)
	return res
}

func main() {
	defer stdout.Flush()
	h := NewHeap()
	for {
		cmd := readString()
		switch cmd {
		case "insert":
			k := readInt()
			h.Insert(k)
		case "extract":
			k := h.ExtractMax()
			println(k)
		case "end":
			return
		default:
			panic("unknown command")
		}
	}
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
