package main

import (
	"fmt"
	"time"
)

func main() {
	const unixTime = 1589570165

	var minutes, seconds int
	fmt.Scanf("%d мин. %d сек.", &minutes, &seconds)
	duration := time.Minute*time.Duration(minutes) + time.Second*time.Duration(seconds)
	bstime := time.Unix(unixTime, 0).UTC()
	nwtm := bstime.Add(duration)
	fmt.Println(nwtm.Format(time.UnixDate))
}
