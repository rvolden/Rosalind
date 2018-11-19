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

func p_dist(seqA, seqB string) (p float64) {
    dist, total := 0.0, 0.0
    for i, _ := range seqA {
        total++
        if seqA[i] != seqB[i] {
            dist++
        }
    }
    return dist/total
}

func main() {
    seqs := fasta_reader()
    d_matrix := []float64{}
    for _, seqA := range seqs {
        for _, seqB := range seqs {
            p := p_dist(seqA.seq, seqB.seq)
            d_matrix = append(d_matrix, p)
        }
    }
    iter := len(seqs)
    for i := 0; i < len(d_matrix); {
        fmt.Println(strings.Trim(fmt.Sprint(d_matrix[i:i+iter]), "[]"))
        i += iter
    }
}
