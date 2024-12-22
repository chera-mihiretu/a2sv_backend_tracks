package main

import "fmt"

type Absr interface {
	Abs() uint32
	Square() uint32
}

type MyInt int
type MyFloat float32

func (m MyInt) Abs() uint32 {
	if m < 0 {
		return uint32(-m)
	} else {
		return uint32(m)
	}
}

func (m MyInt) Square() uint32 {
	var value MyInt = (m * m)
	return uint32(value)
}

func (m *MyFloat) Abs() uint32 {
	if *m < 0 {
		return uint32(-(*m))
	} else {
		return uint32((*m))
	}
}

func (m *MyFloat) Square() uint32 {
	var value MyFloat = ((*m) * (*m))
	return uint32(value)
}

func callThis(t Absr) uint32 {
	return t.Abs()
}

func main() {
	var value Absr = MyInt(-187)

	describe(value)

	var n Absr

	describe(n)

	fmt.Println(value.Abs())

}

func describe(i Absr) {
	fmt.Printf("(%v, %T) \n", i, i)
}
