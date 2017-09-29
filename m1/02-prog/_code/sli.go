package main

func main() {
	var sli []float64
	sli = append(sli, 1.2, 3.4, 5.6)
	println("sli[0]=", sli[0])
	println("sli[1]=", sli[1])
	println("sli[2]=", sli[2])

	sli[2]++

	println("sli[0]=", sli[0])
	println("sli[1]=", sli[1])
	println("sli[2]=", sli[2])
}
