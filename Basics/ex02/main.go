package main

import (
	"fmt"
	"os"
)

func getPower() int {
	return 9000
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	power := getPower()

	fmt.Println("It's over", os.Args[1], power)
}
