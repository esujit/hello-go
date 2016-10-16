package main

import (
	consulapi "github.com/hashicorp/consul/api"
	"fmt"
	"log"
	"time"
)

//var wg sync.WaitGroup

func lockExample() {
	key := "test1/lock"
	config := consulapi.DefaultConfig()
	config.Address = "localhost:8500"
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Println("Error in new client")
		return
	}
	log.Println("Success. Got consul api instance.")
	lock, err := consul.LockKey(key)
	if err != nil {
		log.Println("Error in getting lock key")
		return
	}
	log.Println("Got the lock key... ", key)
	leaderCh, err := lock.Lock(nil)

	if err != nil {
		log.Println("Lock error")
		return
	}
	log.Println("Got the lock...")

	if leaderCh == nil {
		log.Fatal("not leader")
	}
	log.Println("Check the leaderChannel...")

	select {
	case <-leaderCh:
		log.Fatal("Leader channel msg.. may be lost leadership")
	case <-time.After(time.Second):
		log.Fatal("Timed out...")

	}

	log.Println("Someone closed leaderCh...")

	return
}

func lockexample2() {

}

func main() {
	fmt.Println("Locking example...")

	//leaderCh1 := make(chan struct{})
	//var appKey string = "test/lock"
	lockExample()

}
