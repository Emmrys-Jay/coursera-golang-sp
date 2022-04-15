/*
* Race condition is a condition of a program containing functions running concurrently accesses a shared variable,
* performing different operations on the same variable.
*
* Therefore, different interleavings of the goroutine will produce different results.
* A race condition can produce unexpected results, therefore it should be avoided when creating concurrent programs.
 */

/*
* In the program below, a variable n is defined with integer value 1.
*
* The race condition occurs because one goroutine was called within the main goroutine to increment the variable
* n by 2 each. The goroutine prints the value after incrementing it.
*
* But based on the interleaving, the program can increment the value twice before the main goroutine exits.
* In a case of interleaving, the program can print:
*
*	3
*	3
*
*	OR
*
*	3
*	3
*	Incremented more than once!
*
*	OR
*
*	3
*	5
*	Incremented more than once!
*
 */

package main

import (
	"fmt"
)

func addTwo(n *int) {
	*n = *n + 2
	fmt.Println(*n)
}

func main() {
	var n = 1
	go addTwo(&n)
	addTwo(&n)
	fmt.Println(n)
	if n > 3 {
		fmt.Println("Incremented more than once!")
	}
}
