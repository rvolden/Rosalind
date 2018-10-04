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
    pos, length int
}

func (fa *fasta) add_seq(sequence string) {
    fa.seq += sequence
}

func revComp(seq string) string {
    seqRunes := []rune(seq)
    rcomp := ""
    for i := len(seqRunes)-1; i > -1; i-- {
        switch {
        case string(seqRunes[i]) == "A":
            rcomp += "T"
        case string(seqRunes[i]) == "C":
            rcomp += "G"
        case string(seqRunes[i]) == "G":
            rcomp += "C"
        case string(seqRunes[i]) == "T":
            rcomp += "A"
        }
    }
    return rcomp
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

func rev_pal(seq string) bool {
    revc := revComp(seq)
    if revc == seq {
        return true
    } else {
        return false
    }
}

func main() {
    seq := fasta_reader()[0].seq
    lens := []int{4, 5, 6, 7, 8, 9, 10, 11, 12}
    pairs := []pair{}
    for _, l := range lens {
        for i := 0; i < len(seq) - l + 1; i++ {
            subseq := seq[i:i+l]
            if rev_pal(subseq) {
                entry := pair{i+1, l}
                pairs = append(pairs, entry)
            }
        }
    }
    for _, p := range pairs {
        fmt.Println(strings.Trim(fmt.Sprint(p), "{}"))
    }
}
