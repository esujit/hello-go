package main

import "fmt"

func main() {

	mySlice := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
	}

	myOtherSlice := []string{
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice)
}
