package geo2d

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Rect struct {
	Orig          Point
	Width, Height float64
}

type Circle struct {
	Orig Point
	R    float64
}

// Translate returns a new Point translated by dx and dy
func Translate(pt Point, dx float64, dy float64) Point {
	pt.X += dx
	pt.Y += dy
	return pt
}

// STARTSURFACER OMIT

type Surfer interface {
	// Surf returns the surface of a 2D shape
	Surf() float64
}

func (r Rect) Surf() float64 {
	return r.Width * r.Height
}

func (c Circle) Surf() float64 {
	return math.Pi * c.R * c.R
}

// Area returns the area of the given Surfer
func Area(s Surfer) float64 { return s.Surf() }

// ENDSURFACER OMIT

// Perimeter returns the perimeter of the given rectangle
func Perimeter(rect Rect) float64 { return 2 * (rect.Width + rect.Height) }

// STARTSTRINGER OMIT

type Stringer interface {
	// String returns a string reprensation of a value.
	String() string
}

func (p Point) String() string {
	return fmt.Sprintf("Point{X=%v, Y=%v}", p.X, p.Y)
}

func (r Rect) String() string {
	return fmt.Sprintf("Rect{Orig=%v, W=%v, H=%v}", r.Orig, r.Width, r.Height)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle{Orig=%v, R=%v}", c.Orig, c.R)
}

// ENDSTRINGER OMIT
