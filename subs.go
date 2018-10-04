package main

import (
    "fmt"
    "os"
    "bufio"
)

func read_simple() (seqs []string) {
    file, err := os.Open(os.Args[1])
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        seqs = append(seqs, line)
    }
    return
}

func get_pos(og, sub string) (pos []int) {
    for i := 0; i < len(og) - len(sub); i++ {
        if og[i:i+len(sub)] == sub {
            pos = append(pos, i+1)
        }
    }
    return
}

func main() {
    seqs := read_simple()
    og, sub := seqs[0], seqs[1]
    positions := get_pos(og, sub)
    fmt.Println(positions)
}
