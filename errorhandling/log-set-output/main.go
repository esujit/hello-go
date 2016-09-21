package main

import (
	"os"
	"fmt"
	"log"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
		//log.Fatal(err)
		 panic(err)
	}

}
