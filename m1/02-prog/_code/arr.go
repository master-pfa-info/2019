package main

func main() {
	arr := [3]uint{10, 11, 12}
	println("arr[0]=", arr[0])
	println("arr[1]=", arr[1])
	println("arr[2]=", arr[2])

	arr[1] = 42
	println("arr[0]=", arr[0])
	println("arr[1]=", arr[1])
	println("arr[2]=", arr[2])
}
