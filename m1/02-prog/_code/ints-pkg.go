package ints

func Gen(n int) []int {
	var o []int
	for i := 0; i < n; i++ {
		o = append(o, i)
	}
	return o
}

func Sum(vs []int) int {
	sum := 0
	for _, v := range vs {
		sum += v
	}
	return sum
}
