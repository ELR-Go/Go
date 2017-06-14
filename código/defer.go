package main

import "fmt"

func main() {
	fmt.Println("Vamos contar?")

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Acabou")
}
