package main

import "fmt"

const pi = 3.14

func main() {
	const world = "World!!"
	fmt.Println("Hello", world)
	fmt.Println("Happy", pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
