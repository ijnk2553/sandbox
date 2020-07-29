package main

import "fmt"

func area(w, h int) int {
	return w * h
}

func multiply(w, h int) (int, int) {
	return w * 2, h * 2
}

func main() {
	v := area(3, 4)
	fmt.Println(v)
	w, h := multiply(3, 4)
	fmt.Println(w, h)
}
