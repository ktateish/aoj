package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Node struct {
	key         int
	left, right *Node
}

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Find(k int) bool {
	x := t.root
	for x != nil {
		if k < x.key {
			x = x.left
		} else if x.key < k {
			x = x.right
		} else {
			return true
		}
	}
	return false
}

func (t *Tree) Delete(k int) {
	y := &t.root
	x := *y
	for {
		if x == nil {
			return
		}
		if k < x.key {
			y = &x.left
		} else if x.key < k {
			y = &x.right
		} else {
			break
		}
		x = *y
	}
	if x.left == nil && x.right == nil {
		*y = nil
	} else if x.left == nil {
		*y = x.right
	} else if x.right == nil {
		*y = x.left
	} else {
		zy := &x.right
		zx := *zy
		for zx.left != nil {
			zy = &zx.left
			zx = *zy
		}
		x.key = zx.key
		*zy = zx.right
	}
}

func (t *Tree) Insert(k int) {
	y := &t.root
	x := *y
	for x != nil {
		if k < x.key {
			y = &x.left
		} else if x.key < k {
			y = &x.right
		} else {
			// already inserted
			return
		}
		x = *y
	}

	*y = &Node{
		key:   k,
		left:  nil,
		right: nil,
	}
}

const (
	Inorder = iota
	Preorder
)

func (t *Tree) Print(order int) {
	var rec func(*Node)
	rec = func(nd *Node) {
		if nd == nil {
			return
		}
		if order == Preorder {
			printf(" %d", nd.key)
		}
		rec(nd.left)
		if order == Inorder {
			printf(" %d", nd.key)
		}
		rec(nd.right)
	}
	rec(t.root)
	println()
}

func (t *Tree) PrintInorder() {
	t.Print(Inorder)
}

func (t *Tree) PrintPreorder() {
	t.Print(Preorder)
}

func main() {
	defer stdout.Flush()
	m := readInt()
	t := NewTree()
	for i := 0; i < m; i++ {
		cmd := readString()
		switch cmd {
		case "insert":
			x := readInt()
			t.Insert(x)
		case "find":
			x := readInt()
			if t.Find(x) {
				println("yes")
			} else {
				println("no")
			}
		case "delete":
			x := readInt()
			t.Delete(x)
		case "print":
			t.PrintInorder()
			t.PrintPreorder()
		default:
			panic("unknown command")
		}
	}
}

var (
	stdin  *bufio.Reader
	stdout *bufio.Writer
)

func init() {
	stdin = bufio.NewReader(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
}

func scanf(f string, args ...interface{}) (int, error) {
	n, err := fmt.Fscanf(stdin, f, args...)
	for {
		r, _, err := stdin.ReadRune()
		if err != nil || !unicode.IsSpace(r) {
			stdin.UnreadRune()
			break
		}
	}
	return n, err
}

func readInt() int {
	var n int
	scanf("%d", &n)
	return n
}

func readIntSlice(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d", &a[i])
	}
	return a
}

func readLengthAndSlice() []int {
	n := readInt()
	return readIntSlice(n)
}

func readString() string {
	var s string
	scanf("%s", &s)
	return s
}

func printf(f string, args ...interface{}) (int, error) {
	return fmt.Fprintf(stdout, f, args...)
}

func println(args ...interface{}) (int, error) {
	return fmt.Fprintln(stdout, args...)
}
