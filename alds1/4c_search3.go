package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bucket struct {
	val   string
	found bool
}

type Hash struct {
	a []bucket
}

func NewHash(size int) *Hash {
	h := &Hash{
		a: make([]bucket, size),
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

func (h *Hash) Insert(s string) error {
	v := h.hashVal(s)
	for i := v; i < len(h.a); i++ {
		if !h.a[i].found {
			h.a[i].val = s
			h.a[i].found = true
			return nil
		}
	}
	for i := 0; i < v; i++ {
		if !h.a[i].found {
			h.a[i].val = s
			h.a[i].found = true
			return nil
		}
	}
	return fmt.Errorf("hash is full")
}

func (h *Hash) Find(s string) bool {
	v := h.hashVal(s)
	for i := v; i < len(h.a); i++ {
		if !h.a[i].found {
			return false
		} else if h.a[i].val == s {
			return true
		}
	}
	for i := 0; i < v; i++ {
		if !h.a[i].found {
			return false
		} else if h.a[i].val == s {
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
