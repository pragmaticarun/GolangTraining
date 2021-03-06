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
	for i := 65; i <= 65+25; i++ {
		fmt.Printf("%d  %o  %b  %#X  %q \n", i, i, i, i, i)
	}

	i := 20

	for i > 10 {
		fmt.Printf("%d  \n", i)
		i--
	}

	for {
		if i == 100 {
			break
		}
		i++
	}
	a := 10
	b := 12

	if result, _ := whatIf(&a, &b); result == false {
		fmt.Println("Two integers are equal")
	}
}
