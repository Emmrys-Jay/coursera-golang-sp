package main

import "fmt"

func main() {
	var val float64
	fmt.Print("Input a floating point number: ")
	_, err := fmt.Scan(&val)
	if err != nil {
		fmt.Println("Error getting user input: ", err)
	}
	fmt.Println(int(val))
}
