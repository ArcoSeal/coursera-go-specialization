package main

import "fmt"

func main() {
	var input float32
	fmt.Printf("Enter a float: ")
	fmt.Scan(&input)
	fmt.Printf("%.0f\n", input) // this will round, not just truncate
}
