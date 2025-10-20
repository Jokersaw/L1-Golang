package main

import (
	"fmt"
)

func main() {
	var seq = [8]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	var ans = make(map[int][]float64)
	for _, value := range seq {
		key := (int(value) / 10) * 10
		ans[key] = append(ans[key], value)
	}
	fmt.Print(ans)
}
