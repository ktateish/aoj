package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func koch(n int) []complex128 {
	points := []complex128{}
	var rec func(int, complex128, complex128)
	rec = func(n int, p, q complex128) {
		if n <= 0 {
			points = append(points, p)
			return
		}
		v := (q - p) / 3
		s := p + v
		t := s + v*cmplx.Rect(1, math.Pi/3.0)
		u := p + 2*v
		rec(n-1, p, s)
		rec(n-1, s, t)
		rec(n-1, t, u)
		rec(n-1, u, q)
	}
	rec(n, (0 + 0i), (100 + 0i))
	points = append(points, (100 + 0i))
	return points
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	points := koch(n)
	for _, p := range points {
		fmt.Printf("%.6f %.6f\n", real(p), imag(p))
	}
}

type Dlist struct {
	head *DLNode
	tail *DLNode
}

type DLNode struct {
	key  Interface
	prev *DLNode
	next *DLNode
}

type Interface interface {
	Equal(Interface) bool
}

func NewDlist() *Dlist {
	dl := &Dlist{}
	dl.head = &DLNode{}
	dl.tail = dl.head
	dl.head.next = dl.tail
	dl.tail.prev = dl.head

	return dl
}

func (d *Dlist) Insert(x Interface) {
	nd := &DLNode{
		x,
		d.head,
		d.head.next,
	}
	d.head.next = nd
	nd.next.prev = nd
}

func (nd *DLNode) Delete() {
	nd.prev.next = nd.next
	nd.next.prev = nd.prev
}

func (d *Dlist) Delete(x Interface) bool {
	for nd := d.head.next; nd != d.tail; nd = nd.next {
		if nd.key.Equal(x) {
			nd.Delete()
			return true
		}
	}
	return false
}

func (d *Dlist) DeleteFirst() bool {
	if d.head.next == d.tail {
		return false
	}
	d.head.next.Delete()
	return true
}

func (d *Dlist) DeleteLast() bool {
	if d.tail.prev == d.head {
		return false
	}
	d.tail.prev.Delete()
	return true
}

func (d *Dlist) Foreach(f func(Interface)) {
	for nd := d.head.next; nd != d.tail; nd = nd.next {
		f(nd.key)
	}
}
