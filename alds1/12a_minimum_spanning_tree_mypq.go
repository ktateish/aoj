package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type edge struct {
	to, cost int
}

func (e edge) Higher(f Interface) bool {
	fe := f.(edge)
	return e.cost < fe.cost
}

func mst(a [][]int) int {
	n := len(a)
	st := make(map[int]bool)
	q := NewPriorityQueue()
	push := func(to, cost int) {
		e := edge{to, cost}
		q.Push(e)
	}
	pop := func() (int, int) {
		ei, ok := q.Pop()
		if !ok {
			panic("failed to pop")
		}
		e := ei.(edge)
		return e.to, e.cost
	}

	start := 0
	total := 0
	push(start, 0)
	for len(st) != n {
		to, cost := pop()
		if st[to] {
			continue
		}
		st[to] = true
		total += cost
		for i := 0; i < n; i++ {
			c := a[to][i]
			if c == -1 {
				continue
			}
			push(i, c)
		}
	}

	return total
}

func main() {
	defer stdout.Flush()
	n := readInt()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			a[i][j] = readInt()
		}
	}
	println(mst(a))
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

type Interface interface {
	Higher(Interface) bool
}

type PriorityQueue struct {
	a []Interface
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{make([]Interface, 1)}
}

func (h *PriorityQueue) Push(k Interface) {
	parent := func(i int) int {
		return i / 2
	}

	var shiftUp func(int)
	shiftUp = func(i int) {
		p := parent(i)
		if p <= 0 || h.a[p].Higher(h.a[i]) {
			return
		}
		h.a[p], h.a[i] = h.a[i], h.a[p]
		shiftUp(p)
	}

	h.a = append(h.a, k)
	shiftUp(len(h.a) - 1)
}

func (h *PriorityQueue) Pop() (Interface, bool) {
	if h.IsEmpty() {
		return nil, false
	}
	left := func(i int) int {
		return i * 2
	}

	right := func(i int) int {
		return i*2 + 1
	}

	var shiftDown func(int)
	shiftDown = func(i int) {
		l := left(i)
		r := right(i)
		n := len(h.a) - 1
		lesser := -1

		if l <= n && h.a[l].Higher(h.a[i]) {
			lesser = l
		} else {
			lesser = i
		}
		if r <= n && h.a[r].Higher(h.a[lesser]) {
			lesser = r
		}

		if lesser != i {
			h.a[i], h.a[lesser] = h.a[lesser], h.a[i]
			shiftDown(lesser)
		}
	}
	res := h.a[1]
	n := len(h.a) - 1
	h.a[1] = h.a[n]
	h.a = h.a[:n]
	shiftDown(1)
	return res, true
}

func (h *PriorityQueue) IsEmpty() bool {
	return len(h.a) == 1
}
