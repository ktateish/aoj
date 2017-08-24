package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bucket struct {
	val  string
	next *bucket
}

type Hash struct {
	a []*bucket
}

func NewHash(size int) *Hash {
	h := &Hash{
		a: make([]*bucket, size),
	}
	return h
}

func (h *Hash) hashVal(s string) int {
	b := []byte(s)
	val := 0
	for _, x := range b {
		val = (val + int(x)) * 43 % len(h.a)
	}
	return val
}

func (h *Hash) Insert(s string) {
	v := h.hashVal(s)
	b := &bucket{
		val:  s,
		next: h.a[v],
	}
	h.a[v] = b
}

func (h *Hash) Find(s string) bool {
	v := h.hashVal(s)
	for p := h.a[v]; p != nil; p = p.next {
		if p.val == s {
			return true
		}
	}
	return false
}

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	n, _ := strconv.Atoi(r.Text())
	h := NewHash(1000009)
	for i := 0; i < n; i++ {
		r.Scan()
		args := strings.Fields(r.Text())
		switch args[0] {
		case "insert":
			h.Insert(args[1])
		case "find":
			if h.Find(args[1]) {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		default:
			fmt.Fprintf(os.Stderr, "unknown command\n")
			os.Exit(1)
		}
	}
}
