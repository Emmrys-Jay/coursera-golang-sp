package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Input a string: ")
	val := bufio.NewScanner(os.Stdin)
	val.Scan() // use `for scanner.Scan()` to keep reading
	line := val.Text()

	if line[0] == 'i' || line[0] == 'I' {
		if strings.Contains(line, "a") || strings.Contains(line, "A") {
			if line[len(line)-1] == 'n' || line[len(line)-1] == 'N' {
				fmt.Println("Found!")
				return
			}
		}
	}
	fmt.Println("Not Found!")
}
