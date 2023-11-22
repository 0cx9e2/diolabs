package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 101; i++ {
		restoDeTres := i % 3
		restoDeCinco := i % 5
		if i == 0 {
			fmt.Println("Pin Pan!")
		}
		if restoDeTres == 0 {
			fmt.Println("Pin!")
		}
		if restoDeCinco == 0 {
			fmt.Println("Pan!")
		}

	}
}
