package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var input string
	var input_int int
	var numbers []int
	for {
		fmt.Printf("Enter an integer (\"X\" to exit): ")
		fmt.Scan(&input)
		if input == "X" {
			break
		} else {
			input_int, _ = strconv.Atoi(input)
			numbers = append(numbers, input_int)
			sort.Ints(numbers)
			fmt.Println(numbers)
		}
	}
}
