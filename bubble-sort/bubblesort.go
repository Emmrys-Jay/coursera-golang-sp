package main

import "fmt"

var swapped = true

func bubbleSort(numbers []int) {
	for i := 0; i < 9; i++ {
		if swapped == true {
			swapped = false
			for i := 0; i < 9; i++ {
				if numbers[i] > numbers[i+1] {
					swap(numbers, i)
					swapped = true
				}
			}
		} else {
			return
		}
	}
}

func swap(numbers []int, index int) {
	x := numbers[index]
	numbers[index] = numbers[index+1]
	numbers[index+1] = x
}

func main() {
	var numbers []int
	fmt.Println("Bubble Sort Algorithm")
	fmt.Println("Enter 10 integers:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d,: ", i+1)
		var n int
		fmt.Scanf("%d", &n)
		numbers = append(numbers, n)
	}
	bubbleSort(numbers)

	fmt.Println(numbers)
}
