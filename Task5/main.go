package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int)
	N := 5

	timeout := time.After(time.Duration(N) * time.Second)

	go func() {
		i := 1
		for {
			ch <- i
			i++
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case val := <-ch:
			log.Printf("The value is: %d", val)
		case <-timeout:
			log.Println("Timeout!")
			return
		}
	}
}

// Также можно было бы решить через контекст с выставленным таймаутом.
// context.WithTimeout(context.Background(), N*time.Second) и ловить по ctx.Done()
