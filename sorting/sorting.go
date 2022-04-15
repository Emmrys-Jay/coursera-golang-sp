package main

import (
	"fmt"
	"strconv"
)

var numbers []int
var a string
var q bool = true

func intro() {
	fmt.Println("Sort an array with four Goroutines")
	fmt.Println("Input 12 numbers:")

	for i := 0; i < 12; i++ {
		fmt.Printf("%d > ", i+1)
		fmt.Scan(&a)
		s, _ := strconv.Atoi(a)
		numbers = append(numbers, s)
	}
}

func sort(nums []int, c chan []int) {
	if q == true {
		c <- nums
	}
	for i := 0; i < len(nums)-1; i++ {
		for k := 0; k < len(nums)-1; k++ {
			if nums[k] > nums[k+1] {
				c := nums[k]
				nums[k] = nums[k+1]
				nums[k+1] = c
			}
		}
	}
}

func main() {
	intro()
	c := make(chan []int)
	n := int(len(numbers) / 4)
	go sort(numbers[:n], c)
	go sort(numbers[n:2*n], c)
	go sort(numbers[2*n:3*n], c)
	go sort(numbers[3*n:], c)
	arr1 := <-c
	arr2 := <-c
	arr3 := <-c
	arr4 := <-c
	fmt.Printf("array 1: %d \narray 2: %d \narray 3: %d \narray 4: %d \n\n", arr1, arr2, arr3, arr4)

	arr1 = append(arr1, arr2...)
	arr1 = append(arr1, arr3...)
	arr1 = append(arr1, arr4...)
	q = false

	sort(arr1, c)
	numbers = arr1
	fmt.Println("Sorted array: ", numbers)
}
