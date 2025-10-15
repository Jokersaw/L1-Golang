package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)

		for i := 1; i <= 5; i++ {
			ch1 <- i
			fmt.Printf("Sended: %d\n", i)
			time.Sleep(2 * time.Second) // Для наглядности вывода
		}

		fmt.Println("Generator has completed its work")
	}()

	go func() {
		defer close(ch2)

		for val := range ch1 {
			ch2 <- val * val
			fmt.Printf("The number %d is squared\n", val)
		}

		fmt.Println("Handler has completed its work")
	}()

	for res := range ch2 {
		fmt.Printf("Received: %d\n", res)
	}

	fmt.Println("The End")
}
