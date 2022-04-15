package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	Food, Locomotion, Noise string
}

func (a Animal) Eat(aName string) {
	fmt.Printf("%s eats %s\n", aName, a.Food)
}

func (a Animal) Move(aName string) {
	fmt.Printf("%s: %s\n", aName, a.Locomotion)
}

func (a Animal) Speak(aName string) {
	fmt.Printf("%s: %s\n", aName, a.Noise)
}

func main() {
	cow := Animal{
		"grass",
		"walk",
		"moo",
	}

	bird := Animal{
		"worms",
		"fly",
		"peep",
	}

	snake := Animal{
		"mice",
		"slither",
		"hsss",
	}

	fmt.Println("Enter A request in the form: ")
	fmt.Println("\"animal\" \"action\"")
	fmt.Println("Example: cow eat, snake move")
	fmt.Print("\n\n")

	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			txt := scanner.Text()
			animal := strings.Split(txt, " ")[0]
			action := strings.Split(txt, " ")[1]

			switch animal {
			case "cow":
				if action == "eat" {
					cow.Eat(animal)
				} else if action == "move" {
					cow.Move(animal)
				} else if action == "speak" {
					cow.Speak(animal)
				} else {
					fmt.Println("Invalid Input")
				}

			case "bird":
				if action == "eat" {
					bird.Eat(animal)
				} else if action == "move" {
					bird.Move(animal)
				} else if action == "speak" {
					bird.Speak(animal)
				} else {
					fmt.Println("Invalid Input")
				}

			case "snake":
				if action == "eat" {
					snake.Eat(animal)
				} else if action == "move" {
					snake.Move(animal)
				} else if action == "speak" {
					snake.Speak(animal)
				} else {
					fmt.Println("Invalid Input")
				}

			default:
				fmt.Println("Invalid Input")
			}
		}
	}
}
