package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"
)

func stopByFlag() {
	fmt.Println("\nMethod 1: Flag")
	stop := false

	go func() {
		for {
			if stop {
				fmt.Println("Goroutine completed through the flag")
				return
			}
			fmt.Println("Working..")
			time.Sleep(300 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second)
	stop = true
	time.Sleep(500 * time.Millisecond)
}

func stopBySignal() {
	fmt.Println("\nMethod 2: Channel")
	stop := make(chan struct{})

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Goroutine completed through the channel")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	close(stop)
	time.Sleep(500 * time.Millisecond)
}

func stopByContext() {
	fmt.Println("\nMethod 3: Context")
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine completed through the context")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)
}

func stopByGoexit() {
	fmt.Println("\nMethod 4: Runtime.Goexit")

	go func() {
		fmt.Println("Goroutine was launched")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Ending with Goexit()")
		runtime.Goexit()
	}()

	time.Sleep(1 * time.Second)
}

func stopByClosingChannel() {
	fmt.Println("\nMethod 5: Closing Channel")
	jobs := make(chan int)

	go func() {
		for job := range jobs {
			fmt.Println("Job number:", job)
		}
		fmt.Println("Goroutine completed through channel closure")
	}()

	for i := 0; i < 2; i++ {
		jobs <- i
		time.Sleep(300 * time.Millisecond)
	}
	close(jobs)
	time.Sleep(500 * time.Millisecond)
}

func stopByOsExit() {
	fmt.Println("\nMethod 6: Os.Exit")

	go func() {
		defer fmt.Println("Won't be printed due to os.Exit(0)")

		for {
			fmt.Println("Working...")
			time.Sleep(300 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Call os.Exit(0)")
	os.Exit(0)
}

func main() {
	stopByFlag()
	stopBySignal()
	stopByContext()
	stopByGoexit()
	stopByClosingChannel()
	stopByOsExit()
}
