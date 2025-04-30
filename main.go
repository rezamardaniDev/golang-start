package main

import (
	"fmt"
)

func main() {
	lst1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}

	for index, item := range lst1 {
		fmt.Println("index:", index, "item:", item)
	}

	sayHello("reza")
}

func sayHello(name string) {
	fmt.Printf("Hello %s", name)
}
