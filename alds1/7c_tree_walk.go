package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

const (
	Preorder = iota
	Inorder
	Postorder
)

type Node struct {
	Root  bool
	Left  int
	Right int
}

func walk(t []Node, i, order int) {
	if i == -1 {
		return
	}
	if order == Preorder {
		printf(" %d", i)
	}
	walk(t, t[i].Left, order)
	if order == Inorder {
		printf(" %d", i)
	}
	walk(t, t[i].Right, order)
	if order == Postorder {
		printf(" %d", i)
	}
}

func main() {
	defer stdout.Flush()
	n := readInt()
	t := make([]Node, n)
	for i := 0; i < n; i++ {
		t[i].Root = true
	}
	for i := 0; i < n; i++ {
		id := readInt()
		l := readInt()
		r := readInt()
		t[id].Left = l
		t[id].Right = r
		if l != -1 {
			t[l].Root = false
		}
		if r != -1 {
			t[r].Root = false
		}
	}
	root := -1
	for i := 0; i < n; i++ {
		if t[i].Root {
			root = i
			break
		}
	}
	println("Preorder")
	walk(t, root, Preorder)
	println()

	println("Inorder")
	walk(t, root, Inorder)
	println()

	println("Postorder")
	walk(t, root, Postorder)
	println()
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
