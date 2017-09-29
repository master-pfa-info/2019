package main

// START OMIT

func main() {
	for i := 0; i < 20; i++ {
		vs := gen(i)
		println("i=", i, "==>", sum(vs))
	}
}

func gen(n int) []int {
	var out []int
	for i := 0; i < n; i++ {
		out = append(out, i)
	}
	return out
}

func sum(vs []int) int {
	out := 0
	for _, v := range vs {
		if v%2 == 0 {
			out += v
		}
	}
	return out
}

// END OMIT
