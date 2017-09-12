package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func calcDistances(g [][]int, from int) []int {
	type st struct {
		v int
		d int
	}
	ds := make([]int, len(g))
	for i := 0; i < len(g); i++ {
		ds[i] = -1
	}
	q := NewQueue()
	push := func(v, d int) {
		q.Push(st{v, d})
	}
	pop := func() (int, int) {
		vsti, _ := q.Pop()
		vst := vsti.(st)
		return vst.v, vst.d
	}
	empty := func() bool {
		return q.Empty()
	}

	push(from, 0)
	for !empty() {
		v, d := pop()
		if ds[v] != -1 && ds[v] <= d {
			// already reached
			continue
		}
		ds[v] = d
		for _, u := range g[v] {
			push(u, d+1)
		}
	}

	return ds
}

func main() {
	defer stdout.Flush()
	n := readInt()
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		v := readInt() - 1 // input is 1-origin
		k := readInt()
		for j := 0; j < k; j++ {
			u := readInt() - 1
			g[v] = append(g[v], u)
		}
	}
	ds := calcDistances(g, 0)
	for i, d := range ds {
		printf("%d %d\n", i+1, d)
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

type Queue struct {
	a    []interface{}
	head int
	tail int
}

func NewQueue() *Queue {
	q := &Queue{
		a:    make([]interface{}, 4),
		head: 0,
		tail: 0,
	}
	return q
}

func (q *Queue) Capacity() int {
	return len(q.a) - 1
}

func (q *Queue) Length() int {
	if q.head <= q.tail {
		return q.tail - q.head
	} else {
		return len(q.a) - q.head + q.tail
	}
}

func (q *Queue) Empty() bool {
	return q.tail == q.head
}

func (q *Queue) full() bool {
	return (q.tail+1)%len(q.a) == q.head
}

func (q *Queue) Push(p interface{}) {
	if q.full() {
		newa := make([]interface{}, len(q.a)*2)
		if q.head <= q.tail {
			for i := q.head; i < q.tail; i++ {
				newa[i] = q.a[i]
			}
			// q.head and q.tail are reusable
		} else {
			i := q.head
			j := 0
			for ; i < len(q.a); i++ {
				newa[j] = q.a[i]
				j++
			}
			for i = 0; i < q.tail; i++ {
				newa[j] = q.a[i]
				j++
			}
			q.head = 0
			q.tail = j
		}
		q.a = newa
	}
	q.a[q.tail] = p
	q.tail = (q.tail + 1) % len(q.a)
}

func (q *Queue) Pop() (interface{}, bool) {
	if q.Empty() {
		return nil, false
	}
	res := q.a[q.head]
	q.head = (q.head + 1) % len(q.a)
	return res, true
}
