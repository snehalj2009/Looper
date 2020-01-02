package main

import "fmt"

func main() {
	fmt.Println("Starting Looper")
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even number")
		}

	}
}
