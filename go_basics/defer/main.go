package main

import "fmt"

func name() string {
	fmt.Println("Hello")
	return "Chera"

}

func main() {

	defer fmt.Println("World")

	name := name()
	name += " is cool"
}
