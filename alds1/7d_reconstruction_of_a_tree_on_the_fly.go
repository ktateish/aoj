package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func reconstruct(pre, in []int, sp string) {
	if len(in) == 0 {
		return
	}
	i := 0
	for in[i] != pre[0] {
		i++
	}
	reconstruct(pre[1:], in[:i], " ")
	reconstruct(pre[i+1:], in[i+1:], " ")
	printf("%d%s", pre[0], sp)
}

func main() {
	defer stdout.Flush()
	n := readInt()
	preorder := readIntSlice(n)
	inorder := readIntSlice(n)

	reconstruct(preorder, inorder, "")
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
