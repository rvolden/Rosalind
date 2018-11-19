package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func read_simple() (n []string){
    file, err := os.Open(os.Args[1])
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        n = append(n, line)
    }
    return
}

func main() {
    input := read_simple()
    n, _ := strconv.Atoi(input[0])
    b := len(input)
    ans := n-b
    fmt.Println(ans)
}
