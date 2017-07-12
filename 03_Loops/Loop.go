package main

import "fmt"

func main() {
	for i := 65; i <= 65+25; i++ {
		fmt.Printf("%d  %o  %b  %#X  %q \n", i, i, i, i, i)
	}
}
