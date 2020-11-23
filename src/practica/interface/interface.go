package main

import "fmt"

//Printable ..
type Printable interface {
	print()
}

type person struct {
	name string
}

type figure struct {
	name string
}

func (p *person) print()  {
	fmt.Println("[person]", p.name)
}

func (f *figure) print()  {
	fmt.Println("[figure]", f.name)
}

func invokePrint(p Printable)  {
	p.print()
}

func main() {
	p := &person{name: "Juan"}
	f := &figure{name: "Cubo"}

	invokePrint(p)
	invokePrint(f)
}