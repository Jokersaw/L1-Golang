package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int64 = 5
	setOne := true
	pos := 2 // Pos starts from 1

	fmt.Println("Before:", strconv.FormatInt(x, 2))

	if setOne {
		x = x | (1 << (pos - 1))
	} else {
		x = x &^ (1 << (pos - 1))
	}

	fmt.Println("After:", strconv.FormatInt(x, 2))
	fmt.Println("Result:", x)
}
