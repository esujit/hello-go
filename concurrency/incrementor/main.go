package main

import (
	"sync"
	"time"
	"fmt"
	"math/rand"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
}

func incrementor(s string) {
	for i:=0; i<50; i++ {
		mutex.Lock()
		counter++
		time.Sleep(time.Duration(rand.Intn(3))* time.Microsecond)
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()
	}

	wg.Done()
}
