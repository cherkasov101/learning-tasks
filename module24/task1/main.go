package main

import "fmt"

const size = 10

// insertionSort - function for sorting array with 10 integers
func insertionSort(array [size]int) [size]int {
	for i, n := range array {
		j := i - 1
		for j >= 0 && n < array[j] {
			array[j+1] = array[j]
			j = j - 1
		}
		array[j+1] = n
	}
	return array
}

func main() {
	array := [size]int{9, 5, 7, 11, 0, 3, 2, 8, 5, 15}
	array = insertionSort(array)
	fmt.Println(array)
}
