package main

func main() {
	a := 10
	b := 0
	println("a=", a, "b=", b)
	println("add(a,b)=", add(a, b))
	println("mul(a,b)=", mul(a, b))
	println("div(a,b)=", div(a, b))
	println("mul(mul(a,b),div(a,b))=", mul(mul(a, b), div(a, b)))
}

func add(a int, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}
