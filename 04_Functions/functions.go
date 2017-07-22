package main

import "fmt"

func whatIf(a *int, b *int) (bool, int) {
	var g bool
	var err int
	switch {
	case *a/2 == *b/2:
		g = true
		err = 0
	case *a/2 != *b/2:
		g = false
		err = 0
	default:
		err = 1
	}
	return g, err
}

func main() {
	a := 10
	b := 12

	if result, _ := whatIf(&a, &b); result == false {
		fmt.Println("Two integers are equal")
	}
}
