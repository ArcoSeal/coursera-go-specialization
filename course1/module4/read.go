package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var filename string
	var namelist []name

	fmt.Printf("Enter filename: ")
	fmt.Scan(&filename)

	infile, _ := os.Open(filename)

	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		lineSplit := strings.Split(scanner.Text(), " ")
		namelist = append(namelist, name{lineSplit[0], lineSplit[1]})
	}

	for i, v := range namelist {
		fmt.Println("Entry: ", i)
		fmt.Println("First name: ", v.fname)
		fmt.Println("Last name: ", v.lname)
	}
}
