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

var bases = [...]string{"A", "C", "G", "T"}

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

func make_profile(seqs []fasta) map[string][]int {
    profile := map[string][]int {
        "A":{}, "C":{}, "G":{}, "T":{},
    }
    for i := 0; i < len(seqs[0].seq); i++ {
        for _, b := range bases {
            profile[b] = append(profile[b], 0)
        }
        for _, n := range seqs {
            profile[string(n.seq[i])][i]++
        }
    }
    return profile
}

func make_consensus(profile map[string][]int) (consensus string) {
    for i := 0; i < len(profile["A"]); i++ {
        max, which := 0, ""
        for _, b := range bases {
            if profile[b][i] > max {
                m := &max; *m = profile[b][i]
                w := &which; *w = b
            }
        }
        consensus += which
    }
    return consensus
}

func main() {
    seqs := fasta_reader()
    profile := make_profile(seqs)
    cons := make_consensus(profile)
    fmt.Println(cons)
    for _, b := range bases {
        fmt.Printf("%v: %v\n", b, strings.Trim(fmt.Sprint(profile[b]), "[]"))
    }
}
