package main

// factorized import
import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// function return and imports

func main() {
	fmt.Println("Hellow World")

	fmt.Println(myMath())

	fmt.Println(myTime())

	fmt.Println(myArguments("Chera Mihiretu ", 21, " M "))

	fmt.Println(diddiriirsoo(12374))

	fmt.Println(namingReturn())
}

// math random generator

func myMath() int {
	return rand.Intn(10)
}

// current time teller
func myTime() time.Time {
	return time.Now()
}

// function arguments :- remember the type of the variable comes after the name unlike other languages

func myArguments(name string, age int, gender string) string {
	return name + " " + strconv.Itoa(age) + " " + gender
}

// named return

func diddiriirsoo(number int) (tenth, hundredth, thausands int) {
	tenth = number % 10
	hundredth = number % 100 / 10
	thausands = number % 1000 / 100
	return
}

func namingReturn() (x, y, z int) {
	x = 10
	y = 20
	z = 30

	return
}
