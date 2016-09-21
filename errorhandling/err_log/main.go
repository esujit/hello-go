package main

import (
	"os"
	"log"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {

		log.Println("error happened", err)
		//log.Fatal("error happened", err)
		//panic(err)
	}
}
