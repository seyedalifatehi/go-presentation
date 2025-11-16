package main

import "fmt"

func PrintArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	r := Rectangle{Width: 4, Height: 3}
	c := Circle{Radius: 5}

	PrintArea(r)
	PrintArea(c)
}
