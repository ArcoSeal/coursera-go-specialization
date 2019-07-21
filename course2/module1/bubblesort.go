package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(x []int, i int) {
	temp := x[i]
	x[i] = x[i+1]
	x[i+1] = temp
}

func BubbleSort(x []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(x)-1; i++ {
			if x[i] > x[i+1] {
				Swap(x, i)
				swapped = true
			}
		}

	}
}

func main() {
	var inputs []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter integers: ")
	scanner.Scan()
	inputs = strings.Split(scanner.Text(), " ")
	inputsInt := make([]int, len(inputs))
	for i, v := range inputs {
		inputsInt[i], _ = strconv.Atoi(v)
	}

	BubbleSort(inputsInt)

	fmt.Println(inputsInt)
}
