package main

import "fmt"

func main() {
	c := 'a'

	switch {
	case '0' <= c && c <= '9':
		fmt.Printf("%c는 숫자입니다", c)
	case 'a' <= c && c <= 'z':
		fmt.Printf("%c는 소문자 입니다", c)
	case 'A' <= c && c <= 'Z':
		fmt.Printf("%c는 대문자 입니다", c)
	}
}
