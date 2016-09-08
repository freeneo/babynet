package main

import (
	"fmt"
)

type ErrNegative struct {
     err string
}

func (err ErrNegative) Error() string {
    return err.err
}

func Div(x, y float64) (float64, error) {
    if y == 0.0 {
        return 0, ErrNegative{"y can't be 0"}
    } else {
        return x/y, nil
    }
}

func main() {
	fmt.Println(Div(2, 2))
	fmt.Println(Div(-2, 0))
}

