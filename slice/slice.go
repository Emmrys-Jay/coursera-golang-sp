package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	s := make([]int, 3)

	for k := 0; ; k++ {
		var i string
		fmt.Print("Enter an Integer (Enter X to cancel): ")
		fmt.Scan(&i)

		if i == "X" || i == "x" {
			break
		}
		p, err := strconv.Atoi(i)
		if err != nil {
			log.Fatalln("Input is not an integer: ", err)
			continue
		}
		if k > 2 {
			s = append(s, p)
		} else {
			s[k] = p
		}
	}

	for i := 0; i < len(s)-1; i++ {
		for k := 0; k < len(s)-1; k++ {
			if s[k] > s[k+1] {
				c := s[k]
				s[k] = s[k+1]
				s[k+1] = c
			}
		}
	}

	fmt.Println(s)
}
