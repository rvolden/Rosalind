package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_simple() (in []int) {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		a, _ := strconv.Atoi(line)
		in = append(in, a)
	}
	return
}

func heaps(n int, a []int) {
	c := 0
	for {
		if n > 2 {
			heaps(n-1, a)
		}
		if n <= c+1 {
			break
		} else if n%2 == 1 {
			a[0], a[n-1] = a[n-1], a[0]
		} else {
			a[c], a[n-1] = a[n-1], a[c]
		}
		fmt.Println(strings.Trim(fmt.Sprint(a), "[]"))
		c++
	}
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	in := read_simple()[0]
	to_permute := []int{}
	for i := 1; i <= in; i++ {
		to_permute = append(to_permute, i)
	}
	n := len(to_permute)
	fmt.Println(factorial(n))
	fmt.Println(strings.Trim(fmt.Sprint(to_permute), "[]"))
	heaps(n, to_permute)
}
