package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(arr))

	for i := range arr {
		go func(val int) {
			ch <- val * val
		}(arr[i])
	}

	for range arr {
		fmt.Println(<-ch)
	}
}

// В условии не написано про порядок, но если бы он был нужен, то просто передавали бы в канал не только значение, но и индекс
