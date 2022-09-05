package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func main() {
	rand.Seed(time.Now().UnixNano())
	var array [size]int
	for i := range array {
		array[i] = rand.Intn(10 * size)
	}

	var num int
	fmt.Println("Введите число:")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Некорректный ввод.")
	}

	printArray(array)
	fmt.Println("Колличество чисел в массиве после введенного:", countAfter(array, num))
	even, odd := evenAndOdd(array)
	fmt.Printf("Кооличество чётных чисел в массиве: %d\nКолличество нечётных чисел в массиве: %d", even, odd)
}

func printArray(array [size]int) {
	fmt.Println("Числа в массиве:")
	for _, d := range array {
		fmt.Print(d, " ")
	}
	fmt.Print("\n")
}

func countAfter(array [size]int, limit int) (answer int) {
	for i, d := range array {
		if d == limit {
			answer = size - i - 1
			break
		}
	}
	return
}

func evenAndOdd(array [size]int) (even, odd int) {
	for _, d := range array {
		if d%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return
}
