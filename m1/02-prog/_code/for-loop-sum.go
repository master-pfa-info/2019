package main

// START OMIT

func main() {
	for i := 0; i < 20; i++ {
		v := sum(i)
		println("i=", i, "==>", v)
	}
}

func sum(n int) int {
	out := 0
	vs := squares(n)
	for _, v := range vs {
		out += v
	}
	return out
}

func squares(n int) []int {
	var out []int
	for i := 0; i < n; i++ {
		out = append(out, i*i)
	}
	return out
}

// END OMIT
