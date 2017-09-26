package main

func main() {
	vs := []int{0, 2, 4, 6, 8}

	println("loop 1:")
	for i := 0; i < len(vs); i++ {
		println("i=", i, "==>", vs[i])
	}

	println("loop 2:")
	for i := range vs {
		println("i=", i, "==>", vs[i])
	}

	println("loop 3:")
	for i, v := range vs {
		println("i=", i, "==>", v)
	}
}
