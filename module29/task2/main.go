package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		num := 1
		for {
			select {
			case <-shutdown:
				fmt.Println("Выход из программы")
				return
			default:
				fmt.Printf("Квадрат натурального числа %d = %d\n", num, num*num)
				num++
				time.Sleep(time.Second)
			}
		}
	}()

	wg.Wait()
}
