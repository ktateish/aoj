package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parent(i int) int {
	if i == 0 {
		return 0
	} else {
		return (i - 1) / 4
	}
}

func child(i int, r byte) int {
	switch r {
	case 'A':
		return i*4 + 1
	case 'C':
		return i*4 + 2
	case 'G':
		return i*4 + 3
	case 'T':
		return i*4 + 4
	default:
		return -1
	}
}

var (
	H []byte
)

func index(i int) (int, uint) {
	return i / 8, uint(i % 8)
}

func Set(i int) {
	j, offset := index(i)
	H[j] = H[j] | (1 << uint(offset))
}

func Get(i int) bool {
	j, offset := index(i)
	return H[j]&(1<<uint(offset)) != byte(0)
}

func init() {
	H = make([]byte, 2100000)
}

func Insert(s []byte) {
	nd := 0
	for _, b := range s {
		nd = child(nd, b)
	}
	Set(nd)
}

func Find(s []byte) bool {
	nd := 0
	for _, b := range s {
		nd = child(nd, b)
	}
	return Get(nd)
}

func main() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	r.Scan()
	n, _ := strconv.Atoi(r.Text())
	for i := 0; i < n; i++ {
		r.Scan()
		args := strings.Fields(r.Text())
		switch args[0] {
		case "insert":
			Insert([]byte(args[1]))
		case "find":
			if Find([]byte(args[1])) {
				fmt.Fprintln(w, "yes")
			} else {
				fmt.Fprintln(w, "no")
			}
		default:
			fmt.Fprintf(os.Stderr, "unknown command\n")
			os.Exit(1)
		}
	}
	w.Flush()
}
