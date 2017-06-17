package main

import (
	"fmt"
	"time"
)

func main() {
	tique := time.Tick(100 * time.Millisecond)
	bomba := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tique:
			fmt.Println("tique")
		case <-bomba:
			fmt.Println("BOOM!")
			return
		}
	}
}
