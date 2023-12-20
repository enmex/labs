package main

import (
	"fmt"
	"os"
)

/*
На вход программы подается строка I, содержащей малые буквы английского алфавита.
Пусть подстрокой длиной L с началом S является множество непрерывно следующих символов строки.
Например, строка abcab содержит подстроки
длины 1: a, b, c, a, b
длины 2: ab, bc, ca, ab
длины 3: abc, bca, cab
длины 4: abca, bcab
длины 5: abcab.
В строках длины 1 есть два повторяющихся элемента — a и b.
Пусть весом подстрок длины L является произведение максимального количества повторяющихся подстрок этой длины на длину L.
В примере вес длины 1 есть 2 (2 · 1), длины 2 есть 4 (2 · 2), длины 3 — 3 (1 · 3), длины 4 – 4 и длины 5 – 5.
Требуется найти наибольший из всех весов различных длин. (Для приведенного примера это 5)
*/

func getSubstrings(str string, size int) []string {
	substrs := make([]string, 0)
	for i := 0; i + size <= len(str); i++ {
		substrs = append(substrs, str[i:i+size])
	}
	return substrs
}

func getRepeatedNumber(slice []string) int {
	set := make(map[string]int)
	maxRepeatedNumber := 0
	for _, el := range slice {
		set[el] += 1
	}
	for _, count := range set {
		if maxRepeatedNumber < count {
			maxRepeatedNumber = count
		}
	}
	return max(maxRepeatedNumber, 1)
}

func FindMaxSubstrWeight(str string) int {
	maxWeight := -1
	for i := 1; i <= len(str); i++ {
		substrings := getSubstrings(str, i)
		repeatedNumber := getRepeatedNumber(substrings)
		fmt.Printf("substrings of length %d: %s, has %d repeated elements\n", i, substrings, repeatedNumber)
		currentWeight := i * repeatedNumber
		if maxWeight < currentWeight {
			maxWeight = currentWeight
		}
	}
	return maxWeight
}

func main() {
	var I string
	fmt.Fscanf(os.Stdin, "%s", &I)
	fmt.Printf("max string weight: %d\n", FindMaxSubstrWeight(I))
}