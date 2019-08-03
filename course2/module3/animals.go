package main

import "fmt"

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	var animal, request string
	animals := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	fmt.Println("Animals: cow/bird/snake")
	fmt.Println("Requests: eat/move/speak")
	fmt.Println("Enter <animal> <request>:")

	for {
		fmt.Printf("> ")
		fmt.Scan(&animal, &request)

		if request == "eat" {
			animals[animal].Eat()
		} else if request == "move" {
			animals[animal].Move()
		} else if request == "speak" {
			animals[animal].Speak()
		}
	}

}
