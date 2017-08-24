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
	child [4]*node
	term  bool
}

type Hash struct {
	root *node
}

func NewHash() *Hash {
	h := &Hash{&node{}}
	return h
}

func (h *Hash) Insert(s string) {
	nd := h.root
	for i := 0; i < len(s); i++ {
		c := -1
		switch s[i] {
		case 'A':
			c = A
		case 'C':
			c = C
		case 'G':
			c = G
		case 'T':
			c = T
		}
		if nd.child[c] == nil {
			nd.child[c] = &node{}
		}
		nd = nd.child[c]
	}
	nd.term = true
}

func (h *Hash) Find(s string) bool {
	nd := h.root
	for i := 0; i < len(s); i++ {
		c := -1
		switch s[i] {
		case 'A':
			c = A
		case 'C':
			c = C
		case 'G':
			c = G
		case 'T':
			c = T
		}
		if nd.child[c] == nil {
			return false
		}
		nd = nd.child[c]
	}
	return nd.term
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
