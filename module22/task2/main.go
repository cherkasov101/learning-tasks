package main

import (
	"fmt"
)

const size = 12

func makeArray() (array [size]int) {
	fmt.Println("Введите упорядоченный массив:")
	for i := range array {
		if _, err := fmt.Scan(&array[i]); err != nil {
			fmt.Println("Некорректный ввод.")
		}
	}
	return
}

// searchNumber - function for searching first position of entered number in the array
func searchNumber(array [size]int, num int) (answer int) {
	answer = -1
	min := 0
	max := size - 1
	for max >= min {
		middle := (max + min) / 2
		if num == array[middle] {
			answer = middle
			middle--
			for middle >= 0 {
				if array[middle] == num {
					answer = middle
					middle--
				} else {
					break
				}
			}
			break
		} else if num > array[middle] {
			min = middle + 1
		} else {
			max = middle - 1
		}
	}
	return
}

func main() {
	array := makeArray()
	var num int
	fmt.Println("Введите число:")
	if _, err := fmt.Scan(&num); err != nil {
		fmt.Println("Некорректный ввод.")
	}

	index := searchNumber(array, num)
	if index == -1 {
		fmt.Println("Числа нет в массиве.")
		return
	}
	fmt.Println("Индекс первого вхождения заданного числа в массив:", searchNumber(array, num))
}
