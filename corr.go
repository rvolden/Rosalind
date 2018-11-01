package main

import (
    "fmt"
    "os"
    "bufio"
)

func fasta_reader() (seqs []string) {
    file, err := os.Open(os.Args[1])
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if line[0] == 62 {
            seqs = append(seqs, "")
        } else {
            a := &seqs[len(seqs)-1]; *a += line
        }
    }
    return seqs
}

func hamming(seqA, seqB string) (dist int) {
    for i, _ := range seqA {
        if seqA[i] != seqB[i] {
            dist++
        }
    }
    return
}

func revComp(seq string) string {
	seqRunes := []rune(seq)
	rcomp := ""
	for i := len(seqRunes) - 1; i > -1; i-- {
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

func contains(used []int, index int) bool {
    for _, i := range used {
        if index == i {
            return true
        }
    }
    return false
}

func contains_str(used []string, seq string) bool {
    for _, s := range used {
        if seq == s {
            return true
        }
    }
    return false
}

func main() {
    seqs := fasta_reader()
    counts, used := []int{}, []int{}
    for i := 0; i < len(seqs); i++ {
        counts = append(counts, 1)
        for j := 0; j < len(seqs); j++ {
            if i == j {continue}
            if seqs[i] == seqs[j] || seqs[i] == revComp(seqs[j]) {
                if !contains(used, i) {
                    counts[i]++
                    used = append(used, i)
                }
            }
        }
    }

    used_seqs := []string{}
    for i, c := range counts {
        if c == 1 {
            for k, s := range seqs {
                if k == i || counts[k] < 2 {continue}
                if hamming(s, seqs[i]) == 1 && !contains_str(used_seqs, seqs[i]) {
                    fmt.Print(seqs[i], "->", s, "\n")
                    used_seqs = append(used_seqs, seqs[i])
                } else if hamming(revComp(s), seqs[i]) == 1 && !contains_str(used_seqs, seqs[i]) {
                    fmt.Print(seqs[i], "->", revComp(s), "\n")
                    used_seqs = append(used_seqs, seqs[i])
                }
            }
        }
    }
}
