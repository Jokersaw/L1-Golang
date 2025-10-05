package main

import (
	"fmt"
	"strconv"
	"sync"
)

type ConcurrentMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(map[string]int),
	}
}

func (s *ConcurrentMap) Set(key string, val int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = val
}

func main() {
	s := NewConcurrentMap()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				k := "key-" + strconv.Itoa(j)
				s.Set(k, id*1000+j)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("done")
}
