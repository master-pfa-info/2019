package main

func main() {
	var arr [10]int

	if len(arr) > 10 {
		println("array is big")

		if len(arr) > 15 {
			println("array is really big")
		}

	} else if len(arr) == 1 {
		println("array length is just one")

	} else {
		println("array is rather small")
	}

	println("array size=", len(arr))
}
