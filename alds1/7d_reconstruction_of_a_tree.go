package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Node struct {
	Id    int
	Left  *Node
	Right *Node
}

func reconstruct(pre, in []int) *Node {
	if len(in) == 0 {
		return nil
	}
	nd := &Node{}
	nd.Id = pre[0]
	i := 0
	for in[i] != pre[0] {
		i++
	}
	nd.Left = reconstruct(pre[1:], in[:i])
	nd.Right = reconstruct(pre[i+1:], in[i+1:])
	return nd
}

func postorderPrint(nd *Node, sp string) {
	if nd == nil {
		return
	}
	postorderPrint(nd.Left, " ")
	postorderPrint(nd.Right, " ")
	printf("%d%s", nd.Id, sp)
}

func main() {
	defer stdout.Flush()
	n := readInt()
	preorder := readIntSlice(n)
	inorder := readIntSlice(n)

	root := reconstruct(preorder, inorder)

	postorderPrint(root, "")
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
