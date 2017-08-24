package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	A = iota
	C
	G
	T
	TERM
)

type node struct {
	child [5]*node
}

type Hash struct {
	root *node
}

func NewHash() *Hash {
	h := &Hash{nil}
	return h
}

func (h *Hash) Insert(s string) {
	var rec func(nd *node, s string) *node
	rec = func(nd *node, s string) *node {
		if nd == nil {
			nd = &node{}
		}
		if s == "" {
			nd.child[TERM] = &node{}
			return nd
		}
		i := strings.Index("ACGT", s[:1])
		nd.child[i] = rec(nd.child[i], s[1:])
		return nd
	}
	h.root = rec(h.root, s)
}

func (h *Hash) Find(s string) bool {
	var rec func(nd *node, s string) bool
	rec = func(nd *node, s string) bool {
		if nd == nil {
			return false
		}
		if s == "" {
			if nd.child[TERM] != nil {
				return true
			} else {
				return false
			}
		} else {
			i := strings.Index("ACGT", s[:1])
			return rec(nd.child[i], s[1:])
		}
	}
	return rec(h.root, s)
}

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	n, _ := strconv.Atoi(r.Text())
	h := NewHash()
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
