package main

import "fmt"

/*
Задано 2 <= N <= 1000 множеств из 3 <= M <= 10000 элементов,
значения которых могут находиться в диапазоне от -2000000000 до 2000000000.
Значения каждого множества задаются в произвольном порядке.
Гарантируется, что для одного множества все задаваемые значения — различные.
Необходимо определить наибольший размер множества, получаемого при пересечении какой-либо (любой) из пар заданных множеств.
*/

func intersectionSize(a, b []int) int {
	set := make(map[int]bool)
	for _, v := range b {
		set[v] = true
	}
	size := 0
	for _, v := range a {
		if set[v] {
			size++
		}
	}
	return size
}

func MaxIntersectionSize(sets [][]int) int {
	maxIntersectionSize := -1
	for i := 0; i < len(sets); i++ {
		for j := i + 1; j < len(sets); j++ {
			intersectionSize := intersectionSize(sets[i], sets[j])
			fmt.Printf("set#%d and set#%d: %d\n", i, j, intersectionSize)
			if maxIntersectionSize < intersectionSize {
				maxIntersectionSize = intersectionSize
			}
		}
	}
	return maxIntersectionSize
}

func main() {
	sets := [][]int{
		{1, 2, 3, 4, 5},
		{3, 4, 5, 6, 7},
		{5, 6, 7, 8, 9},
	}
	fmt.Println(MaxIntersectionSize(sets))
}