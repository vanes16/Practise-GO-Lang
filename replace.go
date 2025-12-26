package replace

import "strings"

func DNAStrand(dna string) string {
	replacer := strings.NewReplacer(
		"A", "T",
		"T", "A",
		"C", "G",
		"G", "C",
	)

	return replacer.Replace(dna)
}
