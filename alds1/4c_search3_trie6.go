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
)

type node struct {
	a    *node
	c    *node
	g    *node
	t    *node
	term bool
}

type Hash struct {
	root *node
}

func NewHash() *Hash {
	h := &Hash{&node{}}
	return h
}

func (h *Hash) Insert(s string) {
	var rec func(nd *node, s string)
	rec = func(nd *node, s string) {
		if len(s) == 0 {
			nd.term = true
			return
		}
		switch s[0] {
		case 'A':
			if nd.a == nil {
				nd.a = &node{}
			}
			rec(nd.a, s[1:])
		case 'C':
			if nd.c == nil {
				nd.c = &node{}
			}
			rec(nd.c, s[1:])
		case 'G':
			if nd.g == nil {
				nd.g = &node{}
			}
			rec(nd.g, s[1:])
		case 'T':
			if nd.t == nil {
				nd.t = &node{}
			}
			rec(nd.t, s[1:])
		}
	}
	rec(h.root, s)
}

func (h *Hash) Find(s string) bool {
	var rec func(nd *node, s string) bool
	rec = func(nd *node, s string) bool {
		if nd == nil {
			return false
		}
		if len(s) == 0 {
			return nd.term
		} else {
			switch s[0] {
			case 'A':
				return rec(nd.a, s[1:])
			case 'C':
				return rec(nd.c, s[1:])
			case 'G':
				return rec(nd.g, s[1:])
			case 'T':
				return rec(nd.t, s[1:])
			}
		}
		return false
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
