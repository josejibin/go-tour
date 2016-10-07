package main

import "fmt"

func swap(string1, string2 string) (string, string) {
	return string2, string1
}

func main() {
	a, b := "hello", "world"
	fmt.Println(swap(a, b))
}
