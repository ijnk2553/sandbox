package main

import "fmt"

func main() {
	x := 7
	table := [][]int{{1, 5, 9}, {2, 6, 5, 13}, {5, 3, 7, 4}}
LOOP:
	for row, rowValue := range table {
		for col, cloValue := range rowValue {
			if cloValue == x {
				fmt.Printf("found %d(row:%d, col:%d)", x, row, col)
				break LOOP
			}
		}
	}
}
