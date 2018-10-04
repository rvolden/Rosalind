package main

import (
	"fmt"
	"os"
)

func count(seq string) (a, c, g, t int) {
	nucs := []rune(seq)
	for i := 0; i < len(nucs); i++ {
		switch {
		case string(nucs[i]) == "A":
			a++
		case string(nucs[i]) == "C":
			c++
		case string(nucs[i]) == "G":
			g++
		case string(nucs[i]) == "T":
			t++
		}
	}
	return
}

func main() {
	s := os.Args[1]
	a, c, g, t := count(s)
	fmt.Println(a, c, g, t)
}
