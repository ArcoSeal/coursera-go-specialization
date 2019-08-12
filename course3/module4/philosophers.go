package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

// EatRequests are sent by Philopshers.Dine() to the Host
// they contain the Philsopher's ID and a temporary channel for the Host to give (or deny) eating permission on
type EatRequest struct {
	id    int
	ac chan bool
}

type Chopstick struct {
	sync.Mutex
}

// a Philsopher has an ID, a left & right chopstick shared with its neighbours, some flags and a count of the servings it has eaten
type Philosopher struct {
	id             int
	Lstick, Rstick *Chopstick
	isFull         bool
	isEating       bool
	servings       int
}

// Dine() will continuously try and eat until the Philsopher is full (has eaten maxServ times)
// it does this by sending an EatRequest to the Host. if the Host grants permission it will initiate eating
// if the Host denies permission it will try again
// when it is done eating it notifies the Host via another channel
func (p *Philosopher) Dine(rc chan EatRequest, dc chan bool, maxServ int) {
	p.servings = 0
	p.isFull = false
	for p.servings < maxServ {
		fmt.Printf("%d: requesting to eat\n", p.id)
		ac := make(chan bool) // "callback" channel for eating permission
		req := EatRequest{p.id, ac}
		rc <- req // request permission from host
		fmt.Printf("%d: awaiting permission to eat\n", p.id)

		v, ok := <-ac
		if v && ok { // true was sent on ac -> permission was granted
			fmt.Printf("%d: received permission to eat\n", p.id)
			p.isEating = true
			p.Eat() // get chopsticks & eat
			p.isEating = false
			p.servings++
			fmt.Printf("%d: done eating, has had %d servings\n", p.id, p.servings)
			dc <- true // notify host that someone is done eating
		} else if !ok { // ac was closed -> permission was denied
			fmt.Printf("%d: got rejected by host\n", p.id)
		}
		time.Sleep(100 * time.Millisecond) // limit looping speed
	}
	p.isFull = true
	fmt.Printf("%d: is full\n", p.id)
}

func (p *Philosopher) Eat() {
	// acquire chopsticks
	fmt.Printf("%d: Getting chopsticks...\n", p.id)
	p.Lstick.Lock()
	p.Rstick.Lock()

	// eat dem noodles
	fmt.Printf("Starting to eat: %d\n", p.id)
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond) // eat for up to 2 seconds
	fmt.Printf("Finishing eating: %d\n", p.id)

	// release chopsticks
	p.Lstick.Unlock()
	p.Rstick.Unlock()
	fmt.Printf("%d: Put down chopsticks\n", p.id)
}

func Host(philos []*Philosopher, maxEat int, req chan EatRequest, dc chan bool, wg *sync.WaitGroup) {
	var eating int // number of currently eating philosophers
	var alldone bool // will be true when all the philosophers are full

	for !alldone {
		select {
		case er := <-req: // recieved EatRequest from philsopher
			fmt.Printf("Host: got request from %d\n", er.id)
			li, ri := lridx(er.id, len(philos)) // get neighbouring philsophers
			if philos[li].isEating || philos[ri].isEating { // denial
				fmt.Printf("Host: denying %d (neighbour is eating)\n", er.id)
				close(er.ac)
			} else if eating >= maxEat {
				fmt.Printf("Host: denying %d (%d people are already eating)\n", er.id, maxEat)
				close(er.ac)
			} else { // grant
				fmt.Printf("Host: granting permission to %d to eat\n", er.id)
				eating++
				er.ac <- true
			}
		
		case <-dc: // received notification someone is done eating
			eating--
			fmt.Printf("Host: heard someone was done eating: currently eating = %d\n", eating)
		
		default:
			fmt.Printf("Host: No outstanding requests to eat, checking if everyone is full...\n")
			alldone = allfull(philos)
		}
		time.Sleep(100 * time.Millisecond) // limit looping speed
	}
	fmt.Println("Host: everyone is full, ending dinner")
	wg.Done()
}


// get indices to "left" & "right" of an index i, wrapping around at max length n
func lridx(i, n int) (li, ri int) {
	li = i
	ri = i - 1
	if ri < 0 {
		ri = n - 1
	}
	return
}

// check if all philsophers are full of noodles
func allfull(philos []*Philosopher) bool {
	for i := range philos {
		if !philos[i].isFull {
			fmt.Printf("%d is still hungry\n", philos[i].id)
			return false
		}
	}
	return true
}

func main() {
	nPhilos := 5 // number of dining philosophers
	maxServ := 3 // philosopher noodle capacity
	maxEat := 2 // maximum noodle throughput

	var wg sync.WaitGroup
	rc := make(chan EatRequest)
	dc := make(chan bool)

	// slice of pointers to Chopsticks
	chopsticks := make([]*Chopstick, nPhilos)
	for i := range chopsticks {
		chopsticks[i] = &Chopstick{}
	}

	// slice of pointers to Philsophers
	philos := make([]*Philosopher, nPhilos)
	for i := range philos {
		li, ri := lridx(i, nPhilos)
		philos[i] = &Philosopher{i, chopsticks[li], chopsticks[ri], false, false, 0}
	}

	wg.Add(1)
	go Host(philos, maxEat, rc, dc, &wg) // Host starts his shift
	for _, v := range philos { // Philsophers sit down to dinner
		go v.Dine(rc, dc, maxServ)
	}
	wg.Wait()

	fmt.Println("Dinner is over")
}
