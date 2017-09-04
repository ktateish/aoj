package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Node struct {
	Root  bool
	Left  int
	Right int
}

const (
	Preorder = iota
	Inorder
	Postorder
)

func walk(t []Node, i int, order int) {
	stack := []int{i, Preorder}
	push := func(x, o int) {
		stack = append(stack, x, o)
	}
	pop := func() (int, int) {
		x := stack[len(stack)-2]
		o := stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		return x, o
	}
	isEmpty := func() bool {
		return len(stack) == 0
	}

	for !isEmpty() {
		j, o := pop()
		if j == -1 {
			continue
		}
		if o == order {
			printf(" %d", j)
		}
		switch o {
		case Preorder:
			push(j, Inorder)
			push(t[j].Left, Preorder)
		case Inorder:
			push(j, Postorder)
			push(t[j].Right, Preorder)
		case Postorder:
		default:
			panic("unknwon order")
		}
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
