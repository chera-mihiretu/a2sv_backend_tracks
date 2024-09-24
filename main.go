package main

import "fmt"

// this variables has package scope
// meaning they are availeble through out the package
// and must be declared by var
var name, age, gender = "Chera Mihiretu", 21, "M"

var one, two, three int // all of them are going to declared as 0

var isTrue bool // intialized as false

var myString string // going to be initialized to ''

func main() {
	fmt.Println(name, age, gender)
	fmt.Println(one, two, three)
	fmt.Println(isTrue)
	fmt.Println(myString)

	var unval = "This is local variable"

	fmt.Println(unval)
}
