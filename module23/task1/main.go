package main

import "fmt"

func main() {
	even, odd := evenAndOdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Массив чётных чисел:", even)
	fmt.Println("Массив нечётных чисел:", odd)
}

// evenAndOdd - divides an array into even and odd numbers
func evenAndOdd(array ...int) (even, odd []int) {
	for _, i := range array {
		if i%2 == 0 {
			even = append(even, i)
		} else {
			odd = append(odd, i)
		}
	}
	return
}
