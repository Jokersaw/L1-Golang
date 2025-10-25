package main

import "fmt"

func Intersection(a, b []int) []int {
	mapA := make(map[int]bool)
	for _, item := range a {
		mapA[item] = true
	}

	var result []int
	for _, item := range b {
		if mapA[item] {
			result = append(result, item)
		}
	}

	return result
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	intersection := Intersection(A, B)
	fmt.Printf("A = %v\n", A)
	fmt.Printf("B = %v\n", B)
	fmt.Printf("Intersection = %v\n", intersection)
}
