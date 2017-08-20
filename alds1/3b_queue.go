package main

import (
	"bufio"
	"fmt"
	"os"
)

type Proc struct {
	Name string
	Time int
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

func (q *Queue) Pop() interface{} {
	if q.Empty() {
		fatal("queue is empty")
	}
	res := q.a[q.head]
	q.head = (q.head + 1) % len(q.a)
	return res
}

func fatal(s string) {
	fmt.Fprintln(os.Stderr, s)
	os.Exit(1)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, quantum int
	fmt.Fscan(r, &n, &quantum)
	q := NewQueue()
	push := func(p Proc) {
		q.Push(p)
	}
	pop := func() Proc {
		p := q.Pop()
		return p.(Proc)
	}
	for i := 0; i < n; i++ {
		p := Proc{}
		fmt.Fscan(r, &p.Name, &p.Time)
		push(p)
	}

	elapsed := 0
	for !q.Empty() {
		p := pop()
		e := min(quantum, p.Time)
		elapsed += e
		p.Time -= e
		if p.Time == 0 {
			fmt.Println(p.Name, elapsed)
		} else {
			q.Push(p)
		}
	}
}
