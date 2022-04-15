package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type animal struct {
	Food, Locomotion, Noise string
}

type Cow struct {
	Name string
	a    animal
}

type Bird struct {
	Name string
	a    animal
}

type Snake struct {
	Name string
	a    animal
}

func (c Cow) Eat() {
	fmt.Printf("%s\n", c.a.Food)
}

func (c Cow) Move() {
	fmt.Printf("%s\n", c.a.Locomotion)
}

func (c Cow) Speak() {
	fmt.Printf("%s\n", c.a.Noise)
}

func (s Snake) Move() {
	fmt.Printf("%s\n", s.a.Locomotion)
}

func (s Snake) Eat() {
	fmt.Printf("%s\n", s.a.Food)
}

func (s Snake) Speak() {
	fmt.Printf("%s\n", s.a.Noise)
}

func (b Bird) Move() {
	fmt.Printf("%s\n", b.a.Locomotion)
}

func (b Bird) Eat() {
	fmt.Printf("%s\n", b.a.Food)
}

func (b Bird) Speak() {
	fmt.Printf("%s\n", b.a.Noise)
}

func main() {
	cow := Cow{
		"",
		animal{
			"grass",
			"walk",
			"moo",
		},
	}

	bird := Bird{
		"",
		animal{
			"worms",
			"fly",
			"peep",
		},
	}

	snake := Snake{
		"",
		animal{
			"mice",
			"slither",
			"hsss",
		},
	}

	createdAnimals := map[string]string{}
	fmt.Println("Enter A request in the form: ")
	fmt.Println("<req-type> <name> <request/animal-type>")
	fmt.Println("req-type: (newanimal, query)")
	fmt.Println("name: (<any animal name you've added>)")
	fmt.Println("request: (move, eat, speak)")
	fmt.Println("animal-type: (cow, bird, snake)")
	fmt.Println("Enter X to cancel")
	fmt.Print("\n\n")

	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			txt := scanner.Text()
			if txt == "X" {
				return
			}
			first := strings.Split(txt, " ")[0]
			second := strings.Split(txt, " ")[1]
			third := strings.Split(txt, " ")[2]

			switch first {
			case "newanimal":
				if !(third == "bird" || third == "snake" || third == "cow") {
					fmt.Println("Invalid Input")
					continue
				}
				createdAnimals[second] = third //save the animal name and type in a map
				fmt.Println("Created it!")
			case "query":
				val, ok := createdAnimals[second]
				if ok {
					if val == "cow" {
						switch third {
						case "eat":
							fmt.Println(cow.a.Food)
						case "move":
							fmt.Println(cow.a.Locomotion)
						case "speak":
							fmt.Println(cow.a.Noise)
						default:
							fmt.Println("Invalid Input")
						}
					} else if val == "bird" {
						switch third {
						case "eat":
							fmt.Println(bird.a.Food)
						case "move":
							fmt.Println(bird.a.Locomotion)
						case "speak":
							fmt.Println(bird.a.Noise)
						default:
							fmt.Println("Invalid Input")
						}
					} else if val == "snake" {
						switch third {
						case "eat":
							fmt.Println(snake.a.Food)
						case "move":
							fmt.Println(snake.a.Locomotion)
						case "speak":
							fmt.Println(snake.a.Noise)
						default:
							fmt.Println("Invalid Input")
						}
					} else {
						fmt.Println("Invalid Input")
					}
				}
			default:
				fmt.Println("Invalid Input")
			}
		}
	}
}
