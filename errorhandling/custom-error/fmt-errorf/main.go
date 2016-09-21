package main

import (
	"errors"
	"log"
	"fmt"
)

var ErrNoNegativeNum = errors.New("cannot square root of negative num")

func main() {
	fmt.Printf("%T\n", ErrNoNegativeNum)
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("cannot take sqrt of %v", f)
	}

	return 42, nil
}
