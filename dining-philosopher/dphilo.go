package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var philos = []int{1, 2, 3, 4, 5}
var eaten = map[int]int{
	1: 0,
	2: 0,
	3: 0,
	4: 0,
	5: 0,
} //counter map to tell if a philosopher has eaten up to three times
var hostchan = make(chan int, 2)
var decideVar bool = true
var mut sync.RWMutex

func init() {
	rand.Seed(time.Now().UnixMicro())
}

func eat() {
	philonum := rand.Intn(5) + 1
	if eaten[philonum] > 0 {
		return
	}
	for i := 0; i < 3; i++ {
		hostchan <- philos[philonum-1]
		mut.Lock()
		eaten[philonum] += 1
		mut.Unlock()
		fmt.Println("starting to eat ", philonum)
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("finishing eating ", philonum)
		<-hostchan
	}

}

func work() {
	for decideVar {
		go eat()

		mut.Lock()
		for _, v := range eaten {
			if v < 3 {
				decideVar = true
			} else {
				decideVar = false
			}
		}
		mut.Unlock()

	}

}

var deleted int = 0

func main() {
	var values = map[int]int{
		1: 0,
		2: 1,
		3: 2,
		4: 3,
		5: 4,
	}
	//for {
	//val := rand.Intn(5) + 1
	for _, val := range values {
		// if val == v {
		// 	delete(values, k)
		// 	deleted = 1
		// }
		fmt.Printf("value: %d\n", val)
	}
	// if deleted == 1 {
	// 	deleted = 0

	//	continue
	// } else {
	// 	break
	// }
	//}

}
