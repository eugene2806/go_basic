package circle

import "fmt"

func init() {
	fmt.Println("Init from circle package")
}

type Circle struct {
	Radius float64
}

func New(newRadius float64) *Circle {
	return &Circle{Radius: newRadius}
}
