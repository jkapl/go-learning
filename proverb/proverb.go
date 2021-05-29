package proverb

// Proverb converts a string of words into a proverb
func Proverb(rhyme []string) []string {
	var first = "For want of a "
	var last = "And all for the want of a "
	var p = []string{}

	for i, v := range rhyme {
		if i == len(rhyme)-1 {
			var final = last + rhyme[0] + "."
			p = append(p, final)
			break
		}
		var section = first + v + " the " + rhyme[i+1] + " was lost."
		p = append(p, section)
	}

	return p
}
