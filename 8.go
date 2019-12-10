package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var ahora time.Time
	ahora = time.Now()

	fmt.Println(ahora.Year())
	fmt.Println(ahora.Nanosecond())

	rand.Seed(int64(ahora.Nanosecond()))
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(10))
	}

}
