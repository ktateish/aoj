package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type Point struct {
	id   int
	x, y int
}

type Node struct {
	ps            Point
	lower, higher *Node
}

func selector(which string) func(Point) int {
	switch which {
	case "x":
		return func(p Point) int { return p.x }
	case "y":
		return func(p Point) int { return p.y }
	default:
		panic("unkown field")
	}
}

func sorter(which string) func([]Point) {
	switch which {
	case "x":
		return sortByX
	case "y":
		return sortByY
	default:
		panic("unkown field")
	}
}

func nextField(which string) string {
	switch which {
	case "x":
		return "y"
	case "y":
		return "x"
	default:
		panic("unkown field")
	}
}

func Construct(a []Point, field string) *Node {
	switch len(a) {
	case 0:
		return nil
	case 1:
		return &Node{ps: a[0]}
	}

	srt := sorter(field)
	srt(a)

	m := len(a) / 2
	for m > 0 && a[m-1] == a[m] {
		m--
	}
	nf := nextField(field)
	return &Node{
		a[m],
		Construct(a[:m], nf),
		Construct(a[m+1:], nf),
	}
}

func Find(nd *Node, low, high Point, field string) []Point {
	if nd == nil {
		return []Point{}
	}
	sel := selector(field)
	nv := sel(nd.ps)
	lv := sel(low)
	hv := sel(high)
	if nv < lv {
		return Find(nd.higher, low, high, nextField(field))
	} else if hv < nv {
		return Find(nd.lower, low, high, nextField(field))
	} else {
		res := Find(nd.lower, low, high, nextField(field))

		sel2 := selector(nextField(field))
		nv2 := sel2(nd.ps)
		lv2 := sel2(low)
		hv2 := sel2(high)
		if lv2 <= nv2 && nv2 <= hv2 {
			res = append(res, nd.ps)
		}

		res = append(res, Find(nd.higher, low, high, nextField(field))...)
		return res
	}
}

type PointSlice []Point

func (ps PointSlice) Len() int {
	return len(ps)
}
func (ps PointSlice) Less(i, j int) bool {
	return ps[i].id < ps[j].id
}
func (ps PointSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type PointSliceX []Point

func (ps PointSliceX) Len() int {
	return len(ps)
}
func (ps PointSliceX) Less(i, j int) bool {
	return ps[i].x < ps[j].x
}
func (ps PointSliceX) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func sortByX(a []Point) {
	sort.Sort(PointSliceX(a))
}

type PointSliceY []Point

func (ps PointSliceY) Len() int {
	return len(ps)
}
func (ps PointSliceY) Less(i, j int) bool {
	return ps[i].y < ps[j].y
}
func (ps PointSliceY) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func sortByY(a []Point) {
	sort.Sort(PointSliceY(a))
}

func main() {
	defer stdout.Flush()

	n := readInt()
	ps := make([]Point, n)
	for i := 0; i < n; i++ {
		ps[i].id = i
		ps[i].x = readInt()
		ps[i].y = readInt()
	}
	root := Construct(ps, "x")

	q := readInt()
	for i := 0; i < q; i++ {
		var low, high Point
		low.x = readInt()
		high.x = readInt()
		low.y = readInt()
		high.y = readInt()

		ps := Find(root, low, high, "x")

		sort.Sort(PointSlice(ps))
		for _, p := range ps {
			println(p.id)
		}
		println()
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
