package main

import "fmt"
import "strings"

func main() {
	var input string // to avoid precision erorrs with casting floats, we're just going to use strings instead - only the output matters
	fmt.Printf("Enter a float: ")
	fmt.Scan(&input)
	fmt.Println(strings.Split(input, ".")[0])
}
