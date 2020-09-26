// Package strand provides a function for transcribing a dna strand to rna
package strand

// ToRNA takes a dna string and returns its rna string complement
func ToRNA(dna string) string {
	var rna string
	for _, nucleotide := range dna {
		switch nucleotide {
		case 'G':
			rna += "C"
		case 'C':
			rna += "G"
		case 'T':
			rna += "A"
		case 'A':
			rna += "U"
		}
	}

	return rna
}
