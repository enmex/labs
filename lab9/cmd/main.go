package main

import "fmt"

/*
Пусть в каждой ячейке прямоугольной таблицы размером N × M записано число от −109 до 109 .
Будем считать подтаблицей любую часть таблицы, включая целую, образующую прямоугольник, а её весом – сумму всех её чисел.
Необходимо найти вес самой тяжёлой из всех возможных подтаблиц, которые можно построить на основе оригинальной.
*/

func MaxWeight(table [][]int) int {
	rows := len(table)
    cols := len(table[0])

    for i := 1; i < rows; i++ {
        for j := 0; j < cols; j++ {
            table[i][j] += table[i-1][j]
        }
    }

    maxSum := table[0][0]

    for i := 0; i < rows; i++ {
        for j := i; j < rows; j++ {
            currentSum := 0

            for k := 0; k < cols; k++ {
                colSum := table[j][k]
                if i > 0 {
                    colSum -= table[i-1][k]
                }
                currentSum += colSum

                if currentSum > maxSum {
                    maxSum = currentSum
                }

                if currentSum < 0 {
                    currentSum = 0
                }
            }
        }
    }

    return maxSum
}

func main() {
	table := [][]int{
		{-1001, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(MaxWeight(table))
}