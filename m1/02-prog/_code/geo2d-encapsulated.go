package geo2d

type Point struct {
	X, Y float64
}

// STARTTYPES OMIT

type Rect struct{ pts [4]Point }

// NewRectFrom creates a new rectangle from the 4 given corners.
// NewRectFrom will panic if the 4 given points do not form a valid rectangle.
func NewRectFrom(a, b, c, d Point) Rect {
	var r Rect
	// check a, b, c and d do form a valid rectangle
	return r
}

// NewRect creates a new rectangle from an origin point and a width and height.
func NewRect(o Point, w, h float64) Rect {
	return Rect{pts: [4]Point{
		Translate(o, 0, 0),
		Translate(o, w, 0),
		Translate(o, w, h),
		Translate(o, 0, h),
	}}
}

// ENDTYPES OMIT

// Translate returns a new Point translated by dx and dy
func Translate(pt Point, dx float64, dy float64) Point {
	pt.X += dx
	pt.Y += dy
	return pt
}
