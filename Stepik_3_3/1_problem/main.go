package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(num uint) uint {
		numStr := strconv.FormatUint(uint64(num), 10)
		rs := make([]rune, 0, len(numStr))

		for _, r := range numStr {
			if r%2 == 0 {
				rs = append(rs, r)
			}
		}
		numStr = string(rs)

		x, _ := strconv.Atoi(numStr)
		if x == 0 {
			x = 100
		}

		return uint(x)
	}

	fmt.Println(fn(727178))
}
