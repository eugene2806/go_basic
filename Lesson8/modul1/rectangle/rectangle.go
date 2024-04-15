package rectangle

type Rectangle struct {
	A, B  int
	color string
}

func New(newA, newB int, newColor string) *Rectangle {
	return &Rectangle{newA, newB, newColor}
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.A + r.B)
}
