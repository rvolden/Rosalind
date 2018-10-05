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
    s1, s2 := seqs[0].seq, seqs[1].seq
    // transition : A-G C-T
    transitions, transversions := 0.0, 0.0
    for i, _ := range s1 {
        base1, base2 := string(s1[i]), string(s2[i])
        if base1 != base2 {
            if (base1 == "A" && base2 == "G") || (base1 == "G" && base2 == "A") {
                transitions++
            } else if (base1 == "C" && base2 == "T") || (base1 == "T" && base2 == "C") {
                transitions++
            } else {
                transversions++
            }
        }
    }
    fmt.Println(transitions/transversions)
}
