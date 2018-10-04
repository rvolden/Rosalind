package main

import (
	"bufio"
	"fmt"
	"os"
)

func read_simple() (protein string) {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		protein += scanner.Text()
	}
	protein += "."
	return
}

func main() {
	protein := read_simple()
	codons := map[string]int{
		"A": 4, "R": 6, "H": 2, "K": 2,
		"V": 4, "I": 3, "M": 1, "L": 6,
		"F": 2, "Y": 2, "W": 1, "C": 2,
		"G": 4, "P": 4, "Q": 2, "N": 2,
		"T": 4, "S": 6, "D": 2, "E": 2,
		".": 3,
	}
	var combinations int = 1
	for _, a := range protein {
		combinations *= codons[string(a)]
        b := &combinations; *b = *b % 1000000
	}
	fmt.Println(combinations)
}
