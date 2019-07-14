package main

import "fmt"
import "strings"
import "bufio"
import "os"

func main() {
	var input string

	fmt.Printf("Enter string: ")
	scanner := bufio.NewScanner(os.Stdin) // we have to use bufio.Scanner so we can accept strings with an arbitrary number of spaces
	if scanner.Scan() {
		input = scanner.Text()
	}

	input = strings.ToLower(input)

	if strings.HasPrefix(input, "i") && strings.Contains(input, "a") && strings.HasSuffix(input, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
