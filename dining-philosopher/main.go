package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
}

//var hostchan = make(chan int, 4) //host which makes sure only two philosophers can eat at any time
var wg sync.WaitGroup
var mut sync.Mutex
var done []int
var doneBool bool = false

func init() {
	rand.Seed(time.Now().UnixMicro())
}

func (p Philo) eat(philonum int) {

	// TODO: WORK ON HOW VALUE IS RECEIVED FROM HOSTCHAN
	for i := 0; i < 3; i++ {
		//philonum := <-hostchan
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat ", philonum)
		time.Sleep(500 * time.Millisecond)
		fmt.Println("finishing eating ", philonum)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		wg.Done()
		time.Sleep(20 * time.Millisecond)
	}
}

func main() {
	//make chopsticks
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	//define philosophers
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}

	var philonum = map[int]int{
		1: 0,
		2: 1,
		3: 2,
		4: 3,
		5: 4,
	}
	wg.Add(15)
	for _, val := range philonum {
		go philos[val].eat(val)
	}
	wg.Wait()
}
