package main

import "fmt"

func main() {
	var workArray [10]uint8

	for i := range workArray {
		fmt.Scan(&workArray[i])
	}

	var ind1, ind2 uint
	for i := 0; i < 3; i++ {
		fmt.Scan(&ind1, &ind2)
		workArray[ind1], workArray[ind2] = workArray[ind2], workArray[ind1]
	}

	for _, x := range workArray {
		fmt.Print(x, " ")
	}
}
