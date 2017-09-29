package main

func main() {
	x := 10.0
	y := 11.0
	println("x=", x, "y=", y)
	println("fadd(x,y)=", fadd(x, y))
	println("fmul(x,y)=", fmul(x, y))
	println("fma(x,y,z)=", fma(x, y, 2.0))
}

func fadd(x, y float64) float64   { return x + y }
func fmul(x, y float64) float64   { return x * y }
func fdiv(x, y float64) float64   { return x / y }
func fma(x, y, z float64) float64 { return fadd(fmul(x, y), z) }
