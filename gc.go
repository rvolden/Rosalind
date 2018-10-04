package main

import (
    "fmt"
    "os"
    "bufio"
)

type fasta struct {
    header, seq string
}

func (fa *fasta) add_seq(sequence string) {
    fa.seq += sequence
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func fasta_reader() (seqs []fasta) {
    file, err := os.Open(os.Args[1])
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if line[0] == 62 {
            var entry = fasta{string(line[1:]), ""}
            seqs = append(seqs, entry)
        } else {
            seqs[len(seqs)-1].add_seq(line)
        }
    }
    return seqs
}

type gc_c struct {
    name    string
    content float64
}

func gc(seqs []fasta) gc_c {
    max := gc_c{"false", -1.0}
    for _, entry := range seqs {
        a := len(entry.seq)
        var total float64 = float64(a)
        gc := 0.0
        for _, n := range entry.seq {
            if string(n) == "G" || string(n) == "C" {
                gc++
            }
        }
        gc_percent := gc/total*100
        if gc_percent > max.content {
            b := &max; *b = gc_c{entry.header, gc_percent}
        }
    }
    return max
}

func main() {
    seqs := fasta_reader()
    best := gc(seqs)
    fmt.Println(best.name)
    fmt.Println(best.content)
}
