package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Node struct {
	Parent     int
	Depth      int
	Height     int
	Sibling    int
	LeftChild  int
	RightChild int
}

func (n Node) String() string {
	degree := func() int {
		if n.LeftChild == -1 && n.RightChild == -1 {
			return 0
		} else if n.LeftChild == -1 || n.RightChild == -1 {
			return 1
		} else {
			return 2
		}
	}
	ntype := func() string {
		if n.Parent == -1 {
			return "root"
		} else if n.LeftChild == -1 && n.RightChild == -1 {
			return "leaf"
		} else {
			return "internal node"
		}
	}
	return fmt.Sprintf("parent = %d, sibling = %d, degree = %d, depth = %d, height = %d, %s", n.Parent, n.Sibling, degree(), n.Depth, n.Height, ntype())
}

func InitNode(t []Node, i, l, r int) {
	if l == -1 && r == -1 {
		return
	} else if l == -1 {
		t[i].RightChild = r
		t[r].Parent = i
	} else if r == -1 {
		t[i].LeftChild = l
		t[l].Parent = i
	} else {
		t[i].LeftChild = l
		t[i].RightChild = r
		t[l].Parent = i
		t[r].Parent = i
		t[l].Sibling = r
		t[r].Sibling = l
	}
}

func NewTree(n int) []Node {
	t := make([]Node, n)
	for i := 0; i < n; i++ {
		t[i].Parent = -1
		t[i].Sibling = -1
		t[i].LeftChild = -1
		t[i].RightChild = -1
	}
	return t
}

func InitTree(t []Node) {
	var rec func(int, int)
	rec = func(i, depth int) {
		t[i].Depth = depth
		l := t[i].LeftChild
		r := t[i].RightChild
		if l == -1 && r == -1 {
			t[i].Height = 0
		} else if l == -1 {
			rec(r, depth+1)
			t[i].Height = t[r].Height + 1
		} else if r == -1 {
			rec(l, depth+1)
			t[i].Height = t[l].Height + 1
		} else {
			rec(l, depth+1)
			rec(r, depth+1)
			if t[l].Height > t[r].Height {
				t[i].Height = t[l].Height + 1
			} else {
				t[i].Height = t[r].Height + 1
			}
		}

	}
	var root int
	for i := 0; i < len(t); i++ {
		if t[i].Parent == -1 {
			root = i
			break
		}
	}
	rec(root, 0)
}

func main() {
	defer stdout.Flush()
	n := readInt()
	t := NewTree(n)
	for i := 0; i < n; i++ {
		id := readInt()
		l := readInt()
		r := readInt()
		InitNode(t, id, l, r)
	}
	InitTree(t)
	for i := 0; i < n; i++ {
		fmt.Printf("node %d: %s\n", i, t[i].String())
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
