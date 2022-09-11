package main

import "fmt"

// bubbleSort - function for sorting from maximum to minimum
func bubbleSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}

func main() {
	array := []int{9, 5, 7, 11, 0, 3, 2, 8, 5, 15}
	array = bubbleSort(array)
	fmt.Println(array)
}
