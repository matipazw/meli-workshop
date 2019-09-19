package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	var p, _ = strconv.Atoi(os.Args[1])

	for v := 0; 0 <= v && v <= p; v++ {

		if math.Mod(float64(v), 3) == 0 && math.Mod(float64(v), 5) == 0 {
			fmt.Println("Suma:", v+p)
		}
	}
}
