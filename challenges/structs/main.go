package main

import (
	"fmt"
	"log"
)

// Point of x and y axis
type Point struct {
	X int
	Y int
}

// Move method of point type
func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

//Square built up of lenght and location in coordinates
type Square struct {
	Center Point
	Length int
}

// Move method of square type
func (s *Square) Move(dx, dy int) {
	s.Center.Move(dx, dy)
}

//Area method to calculate area of a square
func (s *Square) Area() int {
	return s.Length * s.Length
}

//NewSquare return reference to new square
func NewSquare(x, y, length int) (*Square, error) {

	if length <= 0 {
		return nil, fmt.Errorf("Length cannot be less than 1")
	}

	s := &Square{
		Center: Point{x, y}, // assign point to center of the square
		Length: length,      // assign length to length of the square
	}

	return s, nil

}

func main() {

	s, err := NewSquare(5, 4, 6)

	if err != nil {
		log.Fatalf("Cannot create square")
	}

	fmt.Printf("%+v\n", s)

	s.Move(3, 4)

	fmt.Printf("%+v\n", s)

	fmt.Printf("%+v\n", s.Area())

}
