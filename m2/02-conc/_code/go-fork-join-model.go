package main

// START OMIT
func main() {
	println("in main")
	go sayHello() // "fork" // HL
	println("bye.")
}

func sayHello() {
	println("hello there!")
}

// STOP OMIT
