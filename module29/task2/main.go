package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-shutdown:
				fmt.Println("Выход из программы")
				return
			default:
				var num int
				fmt.Println("Введите число:")
				fmt.Scan(&num)
				fmt.Println("Квадрат:", num*num)
			}
		}
	}()

	wg.Wait()
}
