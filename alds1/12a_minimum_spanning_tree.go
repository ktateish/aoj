package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"strconv"
)

type edge struct {
	to, cost int
}

type Edges []edge

func (e *Edges) Len() int {
	return len(*e)
}

func (e *Edges) Less(i, j int) bool {
	return (*e)[i].cost < (*e)[j].cost
}

func (e *Edges) Swap(i, j int) {
	(*e)[i], (*e)[j] = (*e)[j], (*e)[i]
}

func (e *Edges) Push(x interface{}) {
	xe := x.(edge)
	old := *e
	*e = append(old, xe)
}

func (e *Edges) Pop() interface{} {
	old := []edge(*e)
	n := len(old) - 1
	x := old[n]
	*e = old[:n]
	return x
}

func mst(a [][]int) int {
	n := len(a)
	st := make(map[int]bool)
	es := make([]edge, n)
	h := Edges(es)
	heap.Init(&h)
	push := func(to, cost int) {
		e := edge{to, cost}
		heap.Push(&h, e)
	}
	pop := func() (int, int) {
		ei := heap.Pop(&h)
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
