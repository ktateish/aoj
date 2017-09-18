package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

type edge struct {
	from, to, cost int
}

func (e edge) Higher(fi Interface) bool {
	f := fi.(edge)
	return e.cost < f.cost
}

func sssp(g [][]edge, root int) ([]int, []int) {
	q := NewPriorityQueue()
	push := func(from, to, cost int) {
		e := edge{from, to, cost}
		q.Push(e)
	}
	pop := func() (int, int, int, bool) {
		ei, ok := q.Pop()
		if !ok {
			return -1, -1, -1, false
		}
		e := ei.(edge)
		return e.from, e.to, e.cost, true
	}

	n := len(g)
	visited := make([]bool, n)

	d := make([]int, n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = math.MaxInt32
		p[i] = -1
	}

	push(root, root, 0)
	for found := 0; found < n; {
		pu, u, cost, ok := pop()
		if !ok {
			break
		}
		if visited[u] {
			continue
		}

		found++
		visited[u] = true
		d[u] = cost
		p[u] = pu

		for _, e := range g[u] {
			if !visited[e.to] {
				push(u, e.to, d[u]+e.cost)
			}
		}
	}
	return d, p
}

func main() {
	defer stdout.Flush()
	n := readInt()
	g := make([][]edge, n)
	for i := 0; i < n; i++ {
		u := readInt()
		k := readInt()
		g[u] = make([]edge, k)
		for j := 0; j < k; j++ {
			g[u][j].to = readInt()
			g[u][j].cost = readInt()
		}
	}

	d, _ := sssp(g, 0)
	for i := 0; i < n; i++ {
		printf("%d %d\n", i, d[i])
		//printf("%d %d (%d", i, d[i], i)
		//for j := i; p[j] != -1; j = p[j] {
		//	printf(" <- %d", p[j])
		//}
		//println(")")
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
