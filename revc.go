package main

import "fmt"

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

func main() {
    seq := "ATGCAGTCA"
    rcomp := revComp(seq)
    fmt.Println(rcomp)
}
