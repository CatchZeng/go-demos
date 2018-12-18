package main

import "fmt"

// Add two numbers
func Add(x, y int) int {
	return x + y
}

func main() {
	sum := Add(1, 2)
	fmt.Print(sum)
}
