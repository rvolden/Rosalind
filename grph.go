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

type pair struct {
    head, tail string
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

func contains(pair_list []pair, p pair) bool {
    for _, entry := range pair_list {
        if entry == p {
            return true
        }
    }
    return false
}

func main() {
    seqs := fasta_reader()
    pairs := []pair{}
    for i := 0; i < len(seqs); i++ {
        end := len(seqs[i].seq)
        for j := 0; j < len(seqs); j++ {
            if i == j {
                continue
            }
            rend := len(seqs[j].seq)
            // check first end and last beginning
            if seqs[i].seq[end-3:] == seqs[j].seq[:3] {
                this_pair := pair{seqs[i].header, seqs[j].header}
                if !contains(pairs, this_pair) {
                    pairs = append(pairs, this_pair)
                }
            }
            // check last end and first beginning
            if seqs[i].seq[:3] == seqs[j].seq[rend-3:] {
                this_pair := pair{seqs[j].header, seqs[i].header}
                if !contains(pairs, this_pair) {
                    pairs = append(pairs, this_pair)
                }
            }
        }
    }
    for _, p := range pairs {
        fmt.Println(strings.Trim(fmt.Sprint(p), "{}"))
    }
}
