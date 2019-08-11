package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (c Bird) Eat() {
	fmt.Println("worms")
}

func (c Bird) Move() {
	fmt.Println("fly")
}

func (c Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (c Snake) Eat() {
	fmt.Println("mice")
}

func (c Snake) Move() {
	fmt.Println("slither")
}

func (c Snake) Speak() {
	fmt.Println("hsss")
}

func Create(t string) Animal {
	var a Animal
	switch t {
	case "cow":
		a = Cow{}
	case "bird":
		a = Bird{}
	case "snake":
		a = Snake{}
	}
	return a
}

func Query(a Animal, r string) {
	switch r {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	}
}

func main() {
	var command, arg1, arg2 string
	animals := make(map[string]Animal)

	fmt.Println("Animal types: cow/bird/snake")
	fmt.Println("Request types: eat/move/speak")
	fmt.Println("To create: newanimal <name> <type>")
	fmt.Println("To query:  query <name> <request>")

	for {
		fmt.Printf("> ")
		fmt.Scan(&command, &arg1, &arg2)

		switch command {
		case "newanimal":
			animals[arg1] = Create(arg2)
			fmt.Println("Created it!")
		case "query":
			Query(animals[arg1], arg2)
		}
	}

}
