package main

import "fmt"

type person struct {
	name string
}

func (p *person) print()  {
	fmt.Println("[person]", p.name)
}

func main() {
	p := &person{name: "Juan"}
	p.print()
}