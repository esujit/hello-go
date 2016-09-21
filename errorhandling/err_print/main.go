package main

import (
	"os"
	"fmt"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("error happened", err)
		//log.Println("error happened", err)
		//log.Fatal("error happened", err)
		//panic(err)
	}
}
