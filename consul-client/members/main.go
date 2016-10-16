package main

import (
	consulapi "github.com/hashicorp/consul/api"
	"fmt"
)


func clientConfig() {
	config := consulapi.DefaultConfig()
	config.Address = "localhost:8500"
	consul, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("Error in new client")
		return
	}
	fmt.Println("Success. Got consul api instance.")
	agent := consul.Agent()
	members, err := agent.Members(false)
	if err != nil {
		fmt.Println("Error getting members", err)
		return
	}

	for _, member := range members {
		fmt.Println("consul members ", member.Name)
	}

}

func main() {
	clientConfig()
}
