package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		b1 = PxPyPzE{10, 10, 20, 90}
		b2 = PxPyPzE{10, 10, 30, 20}
		z  = add(b1, b2)
		m  = z.Mass() // HLxxx
	)

	fmt.Printf("invariant mass= %v\n", m)
	fmt.Printf("z= %v\n", z)
}

type PxPyPzE struct {
	px, py, pz, e float64
}

func (p PxPyPzE) Mass() float64 {
	m2 := p.e*p.e - (p.px*p.px + p.py*p.py + p.pz*p.pz)
	return math.Sqrt(m2)
}

func mass(p PxPyPzE) float64 {
	m2 := p.e*p.e - (p.px*p.px + p.py*p.py + p.pz*p.pz)
	return math.Sqrt(m2)
}

func add(p1, p2 PxPyPzE) PxPyPzE {
	return PxPyPzE{p1.px + p2.px, p1.py + p2.py, p1.pz + p2.pz, p1.e + p2.e}
}
