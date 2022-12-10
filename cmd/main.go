package main

import (
	"fmt"
)

func main() {
	run()
}

func run() {
	fmt.Println("Has been executed from Golang")
	entry := []string{"Jack", "John", "Jones"}
	for i, val := range entry {
		fmt.Printf("At position %d, the character %s is present\n", i, val)
	}
}
