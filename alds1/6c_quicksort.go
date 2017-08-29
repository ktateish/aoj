package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Suit  string
	Rank  int
	Order int
}

func quickSort(a []Card, p, r int) {
	partition := func(p, r int) int {
		i := p - 1
		for j := p; j < r; j++ {
			if a[j].Rank <= a[r].Rank {
				i++
				a[i], a[j] = a[j], a[i]
			}
		}
		a[i+1], a[r] = a[r], a[i+1]
		return i + 1
	}

	var rec func(p, r int)
	rec = func(p, r int) {
		if p < r {
			q := partition(p, r)
			rec(p, q-1)
			rec(q+1, r)
		}
	}
	rec(p, r)
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewScanner(os.Stdin)
	r.Scan()

	n, _ := strconv.Atoi(r.Text())

	a := make([]Card, n)
	for i := 0; i < n; i++ {
		r.Scan()
		fs := strings.Fields(r.Text())
		a[i].Suit = fs[0]
		a[i].Rank, _ = strconv.Atoi(fs[1])
		a[i].Order = i
	}

	quickSort(a, 0, n-1)

	stable := true
	prev := Card{"", -1, -1}
	for i := 0; i < n; i++ {
		if prev.Rank == a[i].Rank && prev.Order > a[i].Order {
			stable = false
			break
		}
		prev = a[i]
	}

	if stable {
		fmt.Fprintln(w, "Stable")
	} else {
		fmt.Fprintln(w, "Not stable")
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(w, "%s %d\n", a[i].Suit, a[i].Rank)
	}
	w.Flush()
}
