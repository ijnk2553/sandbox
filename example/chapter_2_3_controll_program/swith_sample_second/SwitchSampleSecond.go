package main

import "fmt"

func main() {
	i := -2
	switch {
	case i < 0:
		fmt.Println("MINUS")
	case i == 0:
		fmt.Println("ZERO")
	case i > 0:
		fmt.Println("PLUS")
	}
}
