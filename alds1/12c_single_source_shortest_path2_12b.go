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
	to, cost int
}

func sssp(g [][]edge, root int) ([]int, []int) {
	n := len(g)
	visited := make([]bool, n)
	d := make([]int, n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = math.MaxInt32
		p[i] = -1
	}

	d[root] = 0
	for {
		mincost := math.MaxInt32
		u := -1
		for i := 0; i < n; i++ {
			if !visited[i] && d[i] < mincost {
				mincost = d[i]
				u = i
			}
		}

		if u == -1 {
			break
		}

		visited[u] = true
		for _, e := range g[u] {
			w := d[u] + e.cost
			if !visited[e.to] && w < d[e.to] {
				d[e.to] = w
				p[e.to] = u
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
