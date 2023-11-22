package main

import "fmt"

func main() {
	for i := 0; i < 101; i++ {
		resto := i % 3
		if resto == 0 {
			fmt.Println("Número divisível por 3:", i)
		}

	}
}
