package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

type Int int

func (i Int) Equal(j Interface) bool {
	return i == j
}

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	n, _ := strconv.Atoi(r.Text())
	d := NewDlist()
	for i := 0; i < n; i++ {
		r.Scan()
		args := strings.Fields(r.Text())
		switch args[0] {
		case "insert":
			x, _ := strconv.Atoi(args[1])
			ix := Int(x)
			d.Insert(ix)
		case "delete":
			x, _ := strconv.Atoi(args[1])
			ix := Int(x)
			d.Delete(ix)
		case "deleteFirst":
			d.DeleteFirst()
		case "deleteLast":
			d.DeleteLast()
		default:
			fmt.Fprintf(os.Stderr, "no such command: %v", args)
		}
	}
	sp := ""
	d.Foreach(func(x Interface) {

		fmt.Printf("%s%d", sp, x.(Int))
		sp = " "
	})
}
