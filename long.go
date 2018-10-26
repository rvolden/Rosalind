package main

import (
    "fmt"
    "os"
    "bufio"
    str "strings"
)

type pair struct {
    head, tail int
}

func fasta_reader() (seqs []string) {
    // This version of the fasta reader just returns the sequences
    file, err := os.Open(os.Args[1])
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if line[0] == 62 {
            var entry string = ""
            seqs = append(seqs, entry)
        } else {
            a := &seqs[len(seqs)-1]; *a += line
        }
    }
    return seqs
}

func contains(pairs []pair, p pair) bool {
    for _, pr := range pairs {
        if pr == p {
            return true
        }
    }
    return false
}

func paths(start int, end int, pairs []pair) (path []pair) {
    for _, p := range pairs {
        if p.head == start {
            path = append(path, p)
            break
        }
    }
    path_len := 1
    for path_len < len(pairs) {
        for _, p := range pairs {
            if path[len(path)-1].tail == end {break}
            if p.head == path[len(path)-1].tail {
                path = append(path, p)
                path_len++
            }
        }
    }
    return path
}

func assemble(seqs []string, path []pair) (assembly string) {
    for i, p := range path {
        head, tail := seqs[p.head], seqs[p.tail]
        if assembly == "" {
            assembly += head
        }
        // head portion
        if i != 0 {
            addHead := str.Index(assembly, head[:(len(head)/2)-1])
            assembly += head[len(assembly) - addHead:]
        }
        // tail portion
        addTail := str.Index(assembly, tail[:(len(tail)/2)-1])
        assembly += tail[len(assembly) - addTail:]
    }
    return assembly
}

func main() {
    seqs := fasta_reader()
    start, end := 0, 0
    pairs := []pair{}

    for i, a := range seqs {
        all, eall := 0, 0
        for j, b := range seqs {
            if i == j {continue}
            if !str.Contains(b, a[:(len(a)/2)]) {
                all++
            }
            if !str.Contains(a, b[:(len(b)/2)]) {
                eall++
            }
            if str.Contains(b, a[:(len(a)/2)]) {
                var entry = pair{j, i}
                if !contains(pairs, entry) {
                    pairs = append(pairs, entry)
                }
            }
            if str.Contains(a, b[:(len(b)/2)]) {
                var entry = pair{i, j}
                if !contains(pairs, entry) {
                    pairs = append(pairs, entry)
                }
            }
        }
        if all == len(seqs) - 1 {
            m := &start; *m = i
        }
        if eall == len(seqs) - 1 {
            n := &end; *n = i
        }
    }

    path := paths(start, end, pairs)
    assembly := assemble(seqs, path)
    fmt.Println(assembly)
}
