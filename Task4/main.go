package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// ПОЯСНЕНИЕ: использую context + waitgroup, таким образом при вызове cancel(),
// который закроет канал Done, я дождусь завершения всех воркеров в группе
// и после завершу программу.
// SIGINT и SIGTERM ловятся в отдельной горутине, ожидая сигнала из sigCh,
// после чего вызывается cancel().
// Считывание данных в канал и ожидание выполнения контекста происходит в главной горутине
// С таким же успехом можно было бы просто завести канал done, закрыть его также в отдельной
// горутине close(done) и ловить в case <-done:
// Но в общем случае, как я понял, использование контекста является общепринятым
// и используется во многих библиотеках, также у него есть дополнительный фунционал
// в виде отмены по таймауту или по дедлайну

// worker reads integers from the channel and prints them to stdout.
// It continues processing until the context is cancelled.
// Once the context is done, the worker prints a stop message and exits.
//
// Parameters:
//
//	ctx - the context used to signal cancellation
//	id  - the worker ID, used for logging output
//	ch  - the channel from which the worker receives integers
//	wg  - WaitGroup used to signal completion of the worker
func worker(ctx context.Context, id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: stopped\n", id)
			return
		case val := <-ch:
			fmt.Printf("Worker %d: %d\n", id, val)
		}
	}
}

// main sets up the worker pool, channel, and signal handling for graceful shutdown.
// It parses the number of workers from the command-line flags, starts the workers,
// and continuously sends incrementing integers into the channel.
// On receiving SIGINT or SIGTERM, it cancels the context, waits for all workers
// to finish, and exits the program.
func main() {
	n := flag.Int("workers", 3, "number of workers")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= *n; i++ {
		wg.Add(1)
		go worker(ctx, i, ch, &wg)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh
		fmt.Println("\nStopping...")
		cancel()
	}()

	input := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Main stopped, waiting for workers...")
			wg.Wait()
			fmt.Println("All workers stopped. Exit.")
			return
		case ch <- input:
			input++
			time.Sleep(500 * time.Millisecond)
		}
	}
}
