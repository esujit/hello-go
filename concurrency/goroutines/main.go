package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func test_concurrency()  {
	numbers := make([]int, 3)
	for idx, v := range numbers {
		wg.Add(1)
		fmt.Println("index ", idx, " value ", v)
		go func(id int) {

			fmt.Println("anonymous function ", id)


		}(idx)
	}
}

func tc_driver() {
	closeCh := make(chan int)
	test_concurrency()
	wg.Wait()

}

func test_channels(closeCh chan int) {
	numbers := make([]int, 3)
	go func(){
		for idx, value := range numbers {
			fmt.Println("go anonymous function loop idx ", idx)
			closeCh <- (value + idx)
		}
		close(closeCh)
	}()
}

func test_channels_driver() {
	syncCh := make(chan int)
	test_channels(syncCh)
	for value := range syncCh {
		fmt.Println("sync channel value ", value)
	}
}

func main() {
	fmt.Println("Testing go routine")
	test_channels_driver()



}
