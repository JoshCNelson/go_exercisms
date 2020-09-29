// Package protein provides functions for protein translation
package protein

var ErrStop error
var ErrInvalidBase error

var codonStrandMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGG": "Tryptophan",
}

// FromCodon takes a codon string and converts it to it's corresponding polypeptide
func FromCodon(codon string) (string, error) {
	return codonStrandMap[codon], nil
}

var rnaStrandMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromRNA takes a rna string and converts it to it's corresponding polypeptide sequence
func FromRNA(rna string) ([]string, error) {
	ret := []string{}

	beginning := 0
	ending := 3

	for beginning < len(rna) {
		strand, ok := rnaStrandMap[rna[beginning:ending]]

		if !ok {
			return ret, nil
		}

		if strand == "STOP" {
			return ret, nil
		}

		ret = append(ret, strand)
		beginning += 3
		ending += 3
	}

	return ret, nil
}
