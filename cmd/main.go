package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("Test_branch_example")
	add(5, 15)
	fmt.Println("Test_branch_child_example_edited")
}

func add(a int, b int) {
	fmt.Printf("a + b = %v", (a + b))
}
