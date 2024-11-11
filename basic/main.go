package main

import "fmt"

func main() {
	x := 10

	if x < 10 {
		fmt.Println("Small number")
	} else {
		fmt.Println("Big number")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i * 2)
	}

	fmt.Println(x)
}
