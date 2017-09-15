package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Interface interface {
	ID() int
}

type content struct {
	x      Interface
	parent int
	rank   int
}

type UnionFindTree struct {
	m map[int]content
}

func NewUnionFindTree() *UnionFindTree {
	return &UnionFindTree{
		make(map[int]content),
	}
}

func (u *UnionFindTree) Add(x Interface) {
	_, ok := u.m[x.ID()]
	if ok {
		return
	}
	u.m[x.ID()] = content{
		x,
		x.ID(),
		0,
	}
}

func (u *UnionFindTree) Find(x Interface) (int, bool) {
	id := x.ID()
	c, ok := u.m[id]
	if !ok {
		return -1, false
	}
	p := c.parent
	if p == id {
		return p, true
	} else {
		c.parent, _ = u.Find(u.m[c.parent].x)
		u.m[id] = c
		return c.parent, true
	}
}

func (u *UnionFindTree) Union(x, y Interface) {
	p, ok := u.Find(x)
	if !ok {
		return
	}
	q, ok := u.Find(y)
	if !ok {
		return
	}
	pc := u.m[p]
	qc := u.m[q]
	if pc.rank > qc.rank {
		qc.parent = p
		u.m[q] = qc
	} else {
		pc.parent = q
		u.m[p] = pc
		if pc.rank == qc.rank {
			qc.rank++
			u.m[q] = qc
		}
	}
}

func (u *UnionFindTree) IsSameSet(x, y Interface) bool {
	p, ok := u.Find(x)
	if !ok {
		return false
	}
	q, ok := u.Find(y)
	if !ok {
		return false
	}
	return p == q
}

func (u *UnionFindTree) Print() {
	for i, c := range u.m {
		printf("%3d: %3d\n", i, c.parent)
	}
}

type id int

func (i id) ID() int {
	return int(i)
}

func main() {
	defer stdout.Flush()
	_ = readInt()
	m := readInt()
	u := NewUnionFindTree()
	add := func(x int) {
		u.Add(id(x))
	}
	union := func(x, y int) {
		u.Union(id(x), id(y))
	}
	same := func(x, y int) bool {
		return u.IsSameSet(id(x), id(y))
	}
	for i := 0; i < m; i++ {
		s := readInt()
		t := readInt()
		add(s)
		add(t)
		union(s, t)
	}
	q := readInt()
	for i := 0; i < q; i++ {
		s := readInt()
		t := readInt()
		if same(s, t) {
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
