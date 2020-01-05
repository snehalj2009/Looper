package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting Looper")
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even number")
		}

	}

	fmt.Println("GoVersion:", os.Getenv("GOVERSION"))
	fmt.Println("Special-Config:", os.Getenv("GOVERSION"))
	
}
