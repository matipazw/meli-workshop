package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	slice := make([]int, 0, len(os.Args)-1)

	for _, v := range os.Args[1:] {

		value, err := strconv.Atoi(v)

		if err != nil {
			fmt.Printf("Value: %s - def: %t\n", v, err)
		} else {
			slice = append(slice, value)
		}
	}

	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))

}
