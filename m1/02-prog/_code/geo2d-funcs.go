package geo2d

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
