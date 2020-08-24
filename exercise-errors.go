package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		return x, nil
	}
	return 0, &MyError{
		time.Now(),
		"Number was negative!!",
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
