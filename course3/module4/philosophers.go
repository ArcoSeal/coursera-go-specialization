package main

import (
	"fmt"
	"sync"
	"time"
)

func mod_floor(x, n int) int { return ((x % n) + n) % n }

type Philosopher struct {
	id             int
	Lstick, Rstick *Chopstick
	isFull         bool
	servings       int
}

func (p Philosopher) dine(req chan chan bool) {
	p.isFull = false
	for p.servings < 3 {
		fmt.Printf("%d requesting to eat\n", p.id)
		approve := make(chan bool)
		req <- approve
		fmt.Printf("%d awaiting permission to eat\n", p.id)
		_ := <-approve
		fmt.Printf("%d received permission to eat\n", p.id)
		p.eat()
	}
	p.isFull = true
	fmt.Printf("%d is full\n")
}

func (p Philosopher) eat() {
	fmt.Printf("%d: Getting chopsticks...\n", p.id)
	p.Lstick.Lock()
	p.Rstick.Lock()

	fmt.Printf("%d: Got chopsticks, eating...\n", p.id)
	time.Sleep(1 * time.Second)

	fmt.Printf("%d: Finished eating, putting down chopsticks\n", p.id)
	p.Lstick.Unlock()
	p.Rstick.Unlock()

	p.servings++
}

type Chopstick struct{ sync.Mutex }

func Waiter(philos *[]Philosopher, req chan chan bool, wg *sync.WaitGroup) {
	var li, ri int
	var eating int
	var alldone bool

	for !alldone {
		select {
		case i := <-req:
			fmt.Printf("Waiter got request from %d\n", (*philos)[i].id)
			ri = mod_floor(i-1, 5)
			li = mod_floor(i+1, 5)
			if !(*philos)[li].isEating && !(*philos)[ri].isEating && eating < 2 {
				fmt.Printf("Waiter granting permission to %d to eat\n", (*philos)[i].id)
				eating++
				*(*philos)[i].GrantEat()
			} else {
				fmt.Printf("%d denied by waiter\n", (*philos)[i].id)
			}
		case <-done:
			fmt.Printf("Waiter heard someone was done eating\n")
			eating--
			// default:
			// 	fmt.Printf("No outstanding requests to eat, waiter is checking if everyone is full...\n")
			// 	for i, _ := range *philos {
			// 		if !(*philos)[i].isFull {
			// 			fmt.Printf("%d is still hungry\n", (*philos)[i].id)
			// 			break
			// 		}
			// 	}
		}
	}
	wg.Done()
}

func main() {
	nPhilos := 5
	chopsticks := make([]Chopstick, nPhilos)
	philos := make([]Philosopher, nPhilos)

	req := make(chan int)
	done := make(chan int)
	var wg sync.WaitGroup

	for i, _ := range philos {
		philos[i] = Philosopher{i, &chopsticks[mod_floor(i+1, 5)], &chopsticks[mod_floor(i-1, 5)], false, false, false, 0}
	}

	wg.Add(1)
	go Waiter(&philos, req, done, &wg)
	for _, v := range philos {
		go v.dine(req, done)
	}
	wg.Wait()
	time.Sleep(30 * time.Second)
}
