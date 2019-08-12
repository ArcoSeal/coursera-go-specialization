package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sortchunk(x []int, wg *sync.WaitGroup) {
	fmt.Printf("Sorting %v\n", x)
	sort.Ints(x)
	wg.Done()
}

func merge(x, y []int) []int {
	i := 0
	j := 0
	k := 0
	z := make([]int, len(x)+len(y))

	for i < len(x) && j < len(y) {
		if x[i] < y[j] {
			z[k] = x[i]
			i++
		} else {
			z[k] = y[j]
			j++
		}
		k++
	}

	for i < len(x) {
		z[k] = x[i]
		i++
		k++
	}

	for j < len(y) {
		z[k] = y[j]
		j++
		k++
	}

	return z
}

func main() {
	var inputs []string
	nchunks := 4

	// Get inputs
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter integers: ")
	scanner.Scan()
	inputs = strings.Split(scanner.Text(), " ")

	arrsize := len(inputs)
	inputsInt := make([]int, arrsize)
	for i, v := range inputs {
		inputsInt[i], _ = strconv.Atoi(v)
	}

	// break input into chunks
	var j int
	chunks := make([][]int, nchunks)
	chunksize := int(math.Floor(float64(arrsize) / float64(nchunks)))
	for i := 0; i < nchunks; i++ {
		if i == (nchunks - 1) {
			j = arrsize
		} else {
			j = (i + 1) * chunksize
		}
		chunks[i] = inputsInt[i*chunksize : j]
	}

	// sort chunks concurrently
	var wg sync.WaitGroup
	for i := range chunks {
		wg.Add(1)
		go sortchunk(chunks[i], &wg)
	}
	wg.Wait()

	// merge chunks one at a time
	var sortedarr []int
	for i := range chunks {
		sortedarr = merge(sortedarr, chunks[i])
	}
	fmt.Printf("Merged: %v\n", sortedarr)

}
