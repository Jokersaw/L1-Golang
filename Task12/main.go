package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, word := range words {
		set[word] = struct{}{}
	}

	fmt.Println("Множество:")
	for word := range set {
		fmt.Println(word)
	}
}
