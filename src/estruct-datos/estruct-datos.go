package main

import "fmt"

func main()  {
	

	//Arrays
	var arr [2]int
	arr[0] = 1
	arr[1] = 2

	for i, v := range arr {
		fmt.Println(i, v)
	}

	//Slices

	//var l []int
	l := make([]int, 10)
	l = append(l, 10)
	fmt.Printf("%p\n", l)
	l = append(l, 100)
	fmt.Printf("%p\n", l)
	l = append(l, 1000)
	fmt.Printf("%p\n", l)

	for i, v := range l {
		fmt.Println(i, v)
	}

	//Maps

	m := make(map[int]string)
	m[0] = "a"
	m[1] = "b"
	m[0] = "c"

	for k,v := range m {
		fmt.Println(k,v)
	}
}