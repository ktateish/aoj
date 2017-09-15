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
	u := &UnionFindTree{}
	u.parent = make([]int, n)
	u.rank = make([]int, n)
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}
	return u
}

func (u *UnionFindTree) Find(x int) int {
	p := u.parent[x]
	if p == x {
		return p
	} else {
		u.parent[x] = u.Find(p)
		return u.parent[x]
	}
}

func (u *UnionFindTree) Union(x, y int) {
	p := u.Find(x)
	q := u.Find(y)
	if u.rank[p] > u.rank[q] {
		u.parent[q] = p
	} else {
		u.parent[p] = q
		if u.rank[p] == u.rank[q] {
			u.rank[q]++
		}
	}
}

func (u *UnionFindTree) IsSameSet(x, y int) bool {
	return u.Find(x) == u.Find(y)
}

func (u *UnionFindTree) Print() {
	for i, x := range u.parent {
		printf("%3d: %3d\n", i, x)
	}
}

func main() {
	defer stdout.Flush()
	n := readInt()
	m := readInt()
	u := NewUnionFindTree(n)
	for i := 0; i < m; i++ {
		s := readInt()
		t := readInt()
		u.Union(s, t)
	}
	q := readInt()
	for i := 0; i < q; i++ {
		s := readInt()
		t := readInt()
		if u.IsSameSet(s, t) {
			println("yes")
		} else {
			println("no")
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
