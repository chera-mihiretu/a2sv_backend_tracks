package main

import "fmt"

func myCall(value *int) {
	*value += 20
}

type Vertex struct {
	x float32
	y float32
}

func main() {
	var to *int
	var value int = 10

	to = &value

	*to += 1

	fmt.Println("Pointers ", value)

	var result int = 10

	myCall(&result)

	fmt.Println(result)

	var vertex Vertex = Vertex{1, 2}
	var another Vertex = Vertex{y: 1, x: 2}
	fmt.Println(vertex.x, vertex.y)

	fmt.Println(another)

}
