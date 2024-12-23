package main

import "fmt"

type AsGenerator []byte

type MyError []byte

func (e MyError) Error() string {
	return fmt.Sprintf("No Space to Write")
}

func (a AsGenerator) Generate() (int, error) {
	if len(a) == 0 {
		return 0, MyError(a)
	}

	for i := range a {
		a[i] = 'A'
	}

	return len(a), nil
}

func main() {
	b := make([]byte, 32)
	myData := AsGenerator(b)
	n, err := myData.Generate()
	if err != nil {
		return
	}
	fmt.Printf("%q", b[:n])
}
