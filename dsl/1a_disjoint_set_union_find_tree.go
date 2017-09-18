package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type UnionFindTree struct {
	parent []int
	rank   []int
}

func NewUnionFindTree(n int) *UnionFindTree {
	uf := &UnionFindTree{
		make([]int, n),
		make([]int, n),
	}

	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}

	return uf
}

func (uf *UnionFindTree) Find(i int) int {
	p := uf.parent[i]
	if p == i {
		return p
	} else {
		uf.parent[i] = uf.Find(p)
		return uf.parent[p]
	}
}

func (uf *UnionFindTree) Unite(x, y int) {
	px := uf.Find(x)
	py := uf.Find(y)

	if px == py {
		return
	}

	if uf.rank[px] < uf.rank[py] {
		uf.parent[px] = py
	} else {
		uf.parent[py] = px
		if uf.rank[px] == uf.rank[py] {
			uf.rank[py]++
		}
	}
}

func (uf *UnionFindTree) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func main() {
	defer stdout.Flush()
	n := readInt()
	q := readInt()
	uf := NewUnionFindTree(n)
	for i := 0; i < q; i++ {
		cmd := readInt()
		x := readInt()
		y := readInt()
		switch cmd {
		case 0:
			uf.Unite(x, y)
		case 1:
			if uf.Same(x, y) {
				println(1)
			} else {
				println(0)
			}
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
