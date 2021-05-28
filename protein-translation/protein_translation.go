package protein

// import "strings"

var codons = map[string]string{
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

func FromCodon(c string) string {
	return codons[c]
}

func FromRNA(r string) []string {
	var codon = []rune{}
	var codons = []string{}
	for i, char := range r {
		codon = append(codon, char)
		if i % 3 == 0 {
			strCodon := string(codon)
			if val, ok := codons[strCodon]; ok {
				codons = append(codons, strCodon)
			}
		}
	}
	return codons
}
