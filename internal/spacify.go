package internal

func Spacify(types []string) []bool {
	spaces := make([]bool, len(types))

	for i := range types {
		if i+1 < len(types) {
			if (types[i] == "pclosed" || types[i] == "word" || types[i] == "punc" || types[i] == "sqclosed" || types[i] == "dqclosed") && (types[i+1] == "word" || types[i+1] == "popened" || types[i+1] == "sqopened" || types[i+1] == "dqopened") {
				spaces[i] = true
			}
		}
	}
	return spaces
}
