package main

import (
	"fmt"
	"strconv"
)

func display(s string) {
	if v, err := strconv.Atoi(s); err != nil {
		fmt.Printf("%s not Integer\n", s)
	} else {
		fmt.Printf("Integer: %d\n", v)
	}
}

func main() {
	display("two")
	display("2")
}
