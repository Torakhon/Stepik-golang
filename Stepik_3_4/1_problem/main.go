package main

import (
	"fmt"
)

func main() {
	value1, value2, operation := readTask()
	if _, ok := value1.(float64); !ok {
		fmt.Printf("value=%v: %T\n", value1, value1)
		return
	}

	if _, ok := value2.(float64); !ok {
		fmt.Printf("value=%v: %T\n", value2, value2)
		return
	}

	switch operation {
	case "+":
		fmt.Printf("%.4f", value1.(float64)+value2.(float64))
	case "-":
		fmt.Printf("%.4f", value1.(float64)-value2.(float64))
	case "*":
		fmt.Printf("%.4f", value1.(float64)*value2.(float64))
	case "/":
		fmt.Printf("%.4f", value1.(float64)/value2.(float64))
	default:
		fmt.Printf("неизвестная операция")
		return
	}
}
