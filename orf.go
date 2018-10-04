package main

import (
	"bufio"
	"fmt"
	"os"
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

type framestart struct {
	start, f int
}

type framestop struct {
	stop, f int
}

func contains(prots []string, prot string) bool {
    for _, p := range prots {
        if p == prot {
            return true
        }
    }
    return false
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

func translate(seq string) (frames []string) {
	dna_to_prot := map[string]string{
		"TTT": "F", "CTT": "L", "ATT": "I", "GGG": "G",
		"GTT": "V", "TTC": "F", "CTC": "L", "ATC": "I",
		"GTC": "V", "TTA": "L", "CTA": "L", "ATA": "I",
		"GTA": "V", "TTG": "L", "CTG": "L", "ATG": "M",
		"GTG": "V", "TCT": "S", "CCT": "P", "ACT": "T",
		"GCT": "A", "TCC": "S", "CCC": "P", "ACC": "T",
		"GCC": "A", "TCA": "S", "CCA": "P", "ACA": "T",
		"GCA": "A", "TCG": "S", "CCG": "P", "ACG": "T",
		"GCG": "A", "TAT": "Y", "CAT": "H", "AAT": "N",
		"GAT": "D", "TAC": "Y", "CAC": "H", "AAC": "N",
		"GAC": "D", "TAA": ".", "CAA": "Q", "AAA": "K",
		"GAA": "E", "TAG": ".", "CAG": "Q", "AAG": "K",
		"GAG": "E", "TGT": "C", "CGT": "R", "AGT": "S",
		"GGT": "G", "TGC": "C", "CGC": "R", "AGC": "S",
		"GGC": "G", "TGA": ".", "CGA": "R", "AGA": "R",
		"GGA": "G", "TGG": "W", "CGG": "R", "AGG": "R",
	}
	starts, stops := []framestart{}, []framestop{}
	for i := 0; i < len(seq)-3; i++ {
		codon := seq[i : i+3]
		if codon == "ATG" || codon == "CAT" {
			mod := 1
			if codon == "CAT" {
				m := &mod
				*m *= -1
			}
			entry := framestart{i+1, ((i+1) % 3) * mod}
			starts = append(starts, entry)
		}
		if dna_to_prot[codon] == "." || dna_to_prot[revComp(codon)] == "." {
			mod := 1
			if dna_to_prot[revComp(codon)] == "." {
				m := &mod
				*m *= -1
			}
			entry := framestop{i+1, ((i+1) % 3) * mod}
			stops = append(stops, entry)
		}
	}
    for _, start := range starts {
        for _, stop := range stops {
            protein := ""
            if start.f > 0 {
                if start.f == stop.f && start.start < stop.stop {
                    for i := start.start-1; i < stop.stop-1; {
                        codon := seq[i : i+3]
                        if dna_to_prot[codon] == "." {
                            break
                        }
                        protein += dna_to_prot[codon]
                        i += 3
                    }
                    if !contains(frames, protein) {
                        frames = append(frames, protein)
                    }
                }
            } else {
                if start.f == stop.f && start.start > stop.stop {
                    for i := stop.stop-1; i < start.start; {
                        codon := revComp(seq[i : i+3])
                        protein = dna_to_prot[codon] + protein
                        i += 3
                    }
					protein := strings.Split(protein, ".")[0]
                    if !contains(frames, protein) {
                        frames = append(frames, protein)
                    }
                }
            }
        }
    }
    return frames
}

func main() {
	seqs := fasta_reader()
	seq := seqs[0].seq
	frames := translate(seq)
	for _, f := range frames {
		fmt.Println(f)
	}
}
