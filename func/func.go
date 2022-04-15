package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acc, i_d, i_v float64) func(t float64) float64 {
	return func(t float64) float64 {
		s := 1/2*(acc*math.Pow(t, 2)) + (i_v * t) + i_d
		return s
	}
}

func main() {
	var acc, i_d, i_v, t float64
	fmt.Println("Displacement Calculator")
	fmt.Print("Enter Acceleration: ")
	fmt.Scan(&acc)
	fmt.Print("Enter Initial Velocity: ")
	fmt.Scan(&i_v)
	fmt.Print("Enter Initial Displacement: ")
	fmt.Scan(&i_d)

	fn := GenDisplaceFn(acc, i_d, i_v)

	fmt.Print("Enter time: ")
	fmt.Scan(&t)
	fmt.Println("Displacement value equals: ", fn(t))
}
