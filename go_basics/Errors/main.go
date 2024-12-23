package main

import (
	"fmt"
	"math"
	"time"
)

type MyFloat float64

func (e MyFloat) Error() string {
	return fmt.Sprintf("What : Number is Negative %v \nWhen : %s", float64(e), time.Now())
}

func Sqrt(v float64) (float64, error) {
	if v < 0 {
		return 0, MyFloat(v)
	}

	return math.Sqrt(v), nil
}

func main() {
	fmt.Println(Sqrt(20))
}
