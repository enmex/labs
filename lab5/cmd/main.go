package main

import (
	"fmt"
	"lab5/ring"
	"math/rand"
	
)

/*
N серых и M белых мышей сидят по кругу.
Кошка ходит по кругу по часовой стрелке и съедает каждую S -тую мышку.
В первый раз счет начинается с серой мышки.
Составить алгоритм определяющий порядок в котором сидели мышки, если через некоторое время осталось K серых и L белых мышей.
*/

func main() {
	r := ring.NewRing[rune]()
	n := 12
	m := 5

	s := 3
	k := 4
	l := 3

	for i := 0; i < k; i++ {
		r.Put('w')
		r.Move(rand.Intn(100))
	}
	for i := 0; i < l; i++ {
		r.Put('g')
		r.Move(rand.Intn(100))
	}
	
	fmt.Println("BEFORE\n" + r.String() + "\n")

	tmpk := k
	tmpl := m
	for tmpk != n {
		r.Put('w')
		r.Move(-s)
		tmpk++
	}
	for tmpl != m {
		r.Put('g')
		r.Move(-s)
		tmpl++
	}
	fmt.Println("AFTER\n" + r.String() + "\n")

	for r.Size() != k + l {
		r.Move(s)
		r.Remove()
		//fmt.Println(r.String())
	}
	fmt.Println("TEST\n" + r.String())
}