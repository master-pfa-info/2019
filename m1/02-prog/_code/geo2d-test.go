package main

import (
	"fmt"

	"uca.fr/geo2d"
)

func main() {
	r1 := geo2d.Rect{geo2d.Point{0, 0}, 10, 20}
	r2 := geo2d.Rect{geo2d.Point{5, 10}, 10, 20}
	fmt.Printf("r1 overlaps w/ r2: %v\n", geo2d.Overlap(r1, r2)) // true

	r2.Orig = geo2d.Point{10, 20}
	fmt.Printf("r1 overlaps w/ r2: %v\n", geo2d.Overlap(r1, r2)) // true

	r2.Orig = geo2d.Point{11, 21}
	fmt.Printf("r1 overlaps w/ r2: %v\n", geo2d.Overlap(r1, r2)) // false

	r2.Orig = geo2d.Point{-1, -1}
	r2.Width = 100
	r2.Height = 200
	fmt.Printf("r1 overlaps w/ r2: %v\n", geo2d.Overlap(r1, r2)) // true
}
