package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var str string
	numChan := make(chan int)
	for {
		var wg sync.WaitGroup
		wg.Add(2)
		fmt.Println("Введите число:")
		fmt.Scan(&str)
		if str == "стоп" {
			break
		}
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Некорректный ввод.")
		}
		go square(num, &wg, numChan)
		go multByTwo(&wg, numChan)
		wg.Wait()
	}
	defer close(numChan)
}

// square - the function for counting the square of a number
func square(num int, wg *sync.WaitGroup, numChan chan int) {
	defer wg.Done()
	num = num * num
	fmt.Println("Квадрат:", num)
	numChan <- num
}

// multByTwo multiplies 2 numbers
func multByTwo(wg *sync.WaitGroup, numChan chan int) {
	defer wg.Done()
	num := (<-numChan) * 2
	fmt.Println("Произведение :", num)
}
