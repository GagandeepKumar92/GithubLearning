package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	add(5, 15)
}

func add(a int, b int) {
	fmt.Printf("a + b = %v", (a + b))
}
