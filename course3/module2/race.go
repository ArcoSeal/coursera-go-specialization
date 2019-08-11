package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	x := 5
	var wg sync.WaitGroup

	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	go func() {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		x = x + 2
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		x = x * 3
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(x)
}
