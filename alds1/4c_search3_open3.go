package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	M   = 1046527
	NIL = -1
)

var (
	H [M]string
)

func getChar(ch rune) int64 {
	switch ch {
	case 'A':
		return 1
	case 'C':
		return 2
	case 'G':
		return 3
	case 'T':
		return 4
	default:
		return 0
	}
}

func getKey(str string) int64 {
	var sum, p int64
	p = 1
	for _, r := range str {
		sum += p * (getChar(r))
		p *= 5
	}
	return sum
}

func h1(key int64) int {
	return int(key % M)
}

func h2(key int64) int {
	return int(1 + (key % (M - 1)))
}

func find(str string) bool {
	key := getKey(str)
	for i := 0; ; i++ {
		h := (h1(key) + i*h2(key)) % M
		if H[h] == str {
			return true
		} else if len(H[h]) == 0 {
			return false
		}
	}
	return false
}

func insert(str string) {
	key := getKey(str)
	for i := 0; ; i++ {
		h := (h1(key) + i*h2(key)) % M
		if H[h] == str {
			return
		} else if len(H[h]) == 0 {
			H[h] = str
			return
		}
	}
	return
}

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	n, _ := strconv.Atoi(r.Text())
	for i := 0; i < n; i++ {
		r.Scan()
		args := strings.Fields(r.Text())
		switch args[0] {
		case "insert":
			insert(args[1])
		case "find":
			if find(args[1]) {
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
