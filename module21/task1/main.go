package main

import "fmt"

func main() {
	fmt.Println("Введите x:")
	var x int16
	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println("Некорректный ввод")
	}

	fmt.Println("Введите y:")
	var y uint8
	_, err = fmt.Scan(&y)
	if err != nil {
		fmt.Println("Некорректный ввод")
	}

	fmt.Println("Введите z:")
	var z float32
	_, err = fmt.Scan(&z)
	if err != nil {
		fmt.Println("Некорректный ввод")
	}

	fmt.Println("Ответ:", calculate(x, y, z))
}

func calculate(x int16, y uint8, z float32) (s float32) {
	s = 2*float32(x) + float32(y)*float32(y) - 3/z
	return
}
