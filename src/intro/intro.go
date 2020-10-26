package main

import "fmt"

func suma(a, b int) (int, error) {
	return a + b, nil
}

func main() {
	fmt.Println(suma(1, 2))
}