package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

type fasta struct {
    header, seq string
}

func (fa *fasta) add_seq(sequence string) {
    fa.seq += sequence
}

func fasta_reader() (seqs []fasta) {
    file, err := os.Open(os.Args[1])
    if err != nil {
        panic(err)
    }
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

func main() {
    seqs := fasta_reader()
    master, subseq := seqs[0].seq, seqs[1].seq
    pos, i := []int{}, 0
    for _, b := range subseq {
        found := false
        for ; i < len(master); {
            if found {
                break
            }
            if string(b) == string(master[i]) {
                pos = append(pos, i+1)
                i++
                found = true
            }
            i++
        }
    }
    fmt.Println(strings.Trim(fmt.Sprint(pos), "[]"))
}
