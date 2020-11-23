package main

import "fmt"

type Person struct {
	Name string
}

func (x *Person) changeName(n string)  { // * a nivel definicion de funcion indica que se requiere una posicion de memoria, no una variable
	(*x).Name = n // * a nivel operador hace referencia al valor que contiene el puntero
	fmt.Printf("%p %p\n", &x, x) // %p direccion de memoria, %v valor que contiene
}

func main() {
	p := Person{"Juan"}
	fmt.Printf("%p %v\n", &p, p)
	(&p).changeName("Pedro") // & devuelve la direccion de memoria de una variable
	fmt.Printf("%p %v\n", &p, p)
}