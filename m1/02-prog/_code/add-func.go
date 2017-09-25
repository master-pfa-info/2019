package main

func main() {
	a := 10
	b := 11
	println("a+b=", add(a, b))
	c := add(a, -b)
	println("a-b=", c)
}

func add(a int, b int) int {
	return a + b
}
