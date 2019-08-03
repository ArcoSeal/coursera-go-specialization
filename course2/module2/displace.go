package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}

	return fn
}

func main() {
	var a, v0, s0, t float64
	fmt.Println("Enter parameters\n===================")
	fmt.Printf("  a: ")
	fmt.Scan(&a)
	fmt.Printf("v_0: ")
	fmt.Scan(&v0)
	fmt.Printf("s_0: ")
	fmt.Scan(&s0)

	displacefn := GenDisplaceFn(a, v0, s0)

	fmt.Printf("Enter time (t): ")
	fmt.Scan(&t)

	fmt.Println(displacefn(t))
}
