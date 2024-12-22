package main

import "fmt"

type chera struct{}

type newInt int

func (c chera) myMethod() string {
	return "Chera is Cool"
}

func (data newInt) theValue() int {
	return int(data)
}

func (v *Vertex) scale(f int64) {
	v.x = v.x * f
	v.y = v.y * f
}

type Vertex struct {
	x, y int64
}

func main() {

	var my_data chera

	fmt.Println(my_data.myMethod())

	var value newInt
	fmt.Println(value.theValue())

	v := Vertex{10, 30}

	v.scale(10)

	fmt.Println(v)

}
