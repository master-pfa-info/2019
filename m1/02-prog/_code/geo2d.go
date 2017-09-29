package geo2d

// START OMIT

type Point struct {
	X, Y float64
}

type Rect struct {
	Orig          Point
	Width, Height float64
}

// Translate returns a new Point translated by dx and dy
func Translate(pt Point, dx float64, dy float64) Point {
	pt.X += dx
	pt.Y += dy
	return pt
}

// Area returns the area of the given rectangle
func Area(rect Rect) float64 { return rect.Width * rect.Height }

// Perimeter returns the perimeter of the given rectangle
func Perimeter(rect Rect) float64 { return 2 * (rect.Width + rect.Height) }

// END OMIT

// STARTFUNCS OMIT

// Contains returns whether the given rectangle contains the given point
func Contains(rect Rect, pt Point) bool {
	xok := rect.Orig.X <= pt.X && pt.X <= rect.Orig.X+rect.Width
	yok := rect.Orig.Y <= pt.Y && pt.Y <= rect.Orig.Y+rect.Height
	return xok && yok
}

// Overlap returns whether 2 rectangles overlap
func Overlap(r1, r2 Rect) bool {
	return overlap(r1, r2) || overlap(r2, r1)
}

func overlap(r1, r2 Rect) bool {
	a := Translate(r2.Orig, 0, 0)
	b := Translate(r2.Orig, 0, r2.Height)
	c := Translate(r2.Orig, r2.Width, r2.Height)
	d := Translate(r2.Orig, r2.Width, 0)

	return Contains(r1, a) || Contains(r1, b) || Contains(r1, c) || Contains(r1, d)
}

// ENDFUNCS OMIT
