package main

import (
	"log"
	"fmt"
)

type NegNumberError struct {
	lat, long string
	err       error
}

func (e *NegNumberError) Error() string {
	return fmt.Sprintf("error %v %v %v", e.lat, e.long, e.err)
}

func main() {
	//fmt.Printf("%T\n", NegNumberError)
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("value %v out of range", f)
		return 0, &NegNumberError{"50.2 N", "20.5 W", nme}
	}

	return 42, nil
}
