package main

import (
	"fmt"
	"strings"
)

type Battery struct {
	capacity int
}

func NewBattery(value string) Battery {
	return Battery{
		capacity: strings.Count(value, "1"),
	}
}

func (b Battery) String() string {
	return fmt.Sprintf("[%10s]", strings.Repeat("X", b.capacity))
}

func main() {
	fmt.Println(NewBattery("1000010011"))
	// [      XXXX]

	var value string
	_, err := fmt.Scan(&value)
	if err != nil {
		panic(err)
	}

	batteryForTest := NewBattery(value)
	fmt.Println(batteryForTest)
}
