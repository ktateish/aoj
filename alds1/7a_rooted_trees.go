package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Node struct {
	Parent int
	Child  []int
	Depth  int
}

func (n Node) Type() string {
	if n.Parent == -1 {
		return "root"
	} else if len(n.Child) == 0 {
		return "leaf"
	} else {
		return "internal node"
	}
}

func (n Node) SprintChild() string {
	s := []string{}
	for _, c := range n.Child {
		s = append(s, fmt.Sprintf("%d", c))
	}
	return strings.Join(s, ", ")
}

func markDepth(t []Node, id int, depth int) {
	t[id].Depth = depth
	for _, c := range t[id].Child {
		markDepth(t, c, depth+1)
	}
}

func main() {
	defer stdout.Flush()
	n := readInt()
	t := make([]Node, n)
	for i := 0; i < n; i++ {
		t[i].Parent = -1
	}
	for i := 0; i < n; i++ {
		id := readInt()
		d := readInt()
		t[id].Child = readIntSlice(d)
		for _, c := range t[id].Child {
			t[c].Parent = id
		}
	}
	p := -1
	for i := 0; i < n; i++ {
		if t[i].Parent == -1 {
			p = i
			break
		}
	}
	markDepth(t, p, 0)
	for i := 0; i < n; i++ {
		printf("node %d: parent = %d, depth = %d, %s, [%s]\n", i, t[i].Parent, t[i].Depth, t[i].Type(), t[i].SprintChild())
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
