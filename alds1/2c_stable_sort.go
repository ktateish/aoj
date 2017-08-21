package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

type Card struct {
	Suit  string
	Value int
}

func NewCard(s string) Card {
	c := Card{
		Suit: s[:1],
	}
	v, err := strconv.Atoi(s[1:])
	if err != nil {
		log.Fatal("NewCard", s)
	}
	c.Value = v
	return c
}

type Cards []Card

func (c Cards) Len() int {
	return len(c)
}

func (c Cards) Less(i, j int) bool {
	return c[i].Value < c[j].Value
}

func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func p(a []Card) {
	sp := ""
	for _, c := range a {
		fmt.Printf("%s%s%d", sp, c.Suit, c.Value)
		sp = " "
	}
	fmt.Println()
}

func sstable(hint []Card, undet []Card) string {
	for i := 0; i < len(hint) && i < len(undet); i++ {
		if hint[i] != undet[i] {
			return "Not stable"
		}
	}
	return "Stable"
}

func main() {
	var n int

	fmt.Scanf("%d", &n)
	a := make([]Card, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scanf("%s", &s)
		a[i] = NewCard(s)
	}
	b := make([]Card, n)
	copy(b, a)
	bubbleSort(Cards(a))
	p(a)
	fmt.Println("Stable")
	selectionSort(Cards(b))
	p(b)
	fmt.Println(sstable(a, b))
}

func bubbleSort(s sort.Interface) int {
	nswap := 0
	flag := true
	for i := 0; flag; i++ {
		flag = false
		for j := s.Len() - 1; j >= i+1; j-- {
			if s.Less(j, j-1) {
				s.Swap(j, j-1)
				nswap++
				flag = true
			}
		}
	}
	return nswap
}

func selectionSort(s sort.Interface) int {
	nswap := 0
	for i := 0; i < s.Len(); i++ {
		minj := i
		for j := i; j < s.Len(); j++ {
			if s.Less(j, minj) {
				minj = j
			}
		}
		if minj != i {
			s.Swap(minj, i)
			nswap++
		}
	}
	return nswap
}
