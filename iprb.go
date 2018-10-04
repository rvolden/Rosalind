package main

import "fmt"

func main() {
    var k, m, n float64 = 15, 25, 22
    N := k + m + n
    prob := 1 - 1/N/(N-1)*(n*(n-1) + n*m + m*(m-1)/4)
    fmt.Println(prob)
}
