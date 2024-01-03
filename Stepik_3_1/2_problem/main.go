package main

import "fmt"

func main() {

	groupCity := map[int][]string{
		10:   {"A_10", "B_10", "C_10"},
		100:  {"A_100", "B_100", "C_100"},
		1000: {"A_1000", "B_1000", "C_1000"},
	}
	fmt.Println(groupCity)

	cityPopulation := map[string]int{
		"A_10": 10, "B_10": 15, "C_10": 20,
		"A_100": 100, "B_100": 200, "C_100": 400,
		"A_1000": 1001, "B_1000": 2000, "C_1000": 3000,
	}

	for _, s := range groupCity[10] {
		delete(cityPopulation, s)
	}
	for _, s := range groupCity[1000] {
		delete(cityPopulation, s)
	}
	fmt.Println(cityPopulation)
}
