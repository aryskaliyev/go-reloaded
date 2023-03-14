package internal

import "regexp"

// Rewrite regex
// Add <squote>, <dquote>
// Proces _quote types accordingly

func Split(str string) ([]string, []string) {
	re := regexp.MustCompile(`(?P<bin>[\(][\s]*[b][i][n][\s]*[,]{0,1}[\s]*[0-9]*[\s]*[\)])|(?P<cap>[\(][\s]*[c][a][p][\s]*[,]{0,1}[\s]*[0-9]*[\s]*[\)])|(?P<up>[\(][\s]*[u][p][\s]*[,]{0,1}[\s]*[0-9]*[\s]*[\)])|(?P<low>[\(][\s]*[l][o][w][\s]*[,]{0,1}[\s]*[0-9]*[\s]*[\)])|(?P<hex>[\(][\s]*[h][e][x][\s]*[,]{0,1}[\s]*[0-9]*[\s]*[\)])|(?P<punc>[\.,\!\?:;]+)|(?P<word>[a-zA-Z0-9]+['-]{0,1}[a-zA-Z0-9]*)|(?P<popen>[\(\[\{])|(?P<pclose>[\)\]\}])|(?P<sqopen>['])|(?P<dqopen>["])`)

	arr := re.FindAllString(str, -1)
	types := make([]string, len(arr))

	punc := `^(?P<punc>[\.,\!\?:;]+$)`
	bin := `^(?P<bin>[\(][\s]*[b][i][n][\s]*[,]{0,1}[\s]*[0-9]*[\s]*[\)])$`
	cap := `^(?P<cap>[\(][\s]*[c][a][p][\s]*[,]{0,1}[\s]*[0-9]*[\s]*[\)])$`
	up := `^(?P<up>[\(][\s]*[u][p][\s]*[,]{0,1}[\s]*[0-9]*[\s]*[\)])$`
	low := `^(?P<low>[\(][\s]*[l][o][w][\s]*[,]{0,1}[\s]*[0-9]*[\s]*[\)])$`
	hex := `^(?P<hex>[\(][\s]*[h][e][x][\s]*[,]{0,1}[\s]*[0-9]*[\s]*[\)])$`
	word := `^(?P<word>[a-zA-Z0-9]+['-]{0,1}[a-zA-Z0-9]*)$`
	popen := `^(?P<popen>[\(\[\{])$`
	pclose := `^(?P<pclosed>[\)\]\}])$`
	sqopen := `^(?P<sqopen>['])$`
	dqopen := `^(?P<dqopen>["])$`

	sqopened := false
	sqclosed := false

	dqopened := false
	dqclosed := false

	for i := range arr {
		if regexp.MustCompile(punc).MatchString(arr[i]) {
			types[i] = "punc"
		} else if regexp.MustCompile(bin).MatchString(arr[i]) {
			types[i] = "bin"
			numb := regexp.MustCompile(`[0-9]+`).FindString(arr[i])
			if numb == "" {
				numb = "1"
			}
			arr[i] = numb
		} else if regexp.MustCompile(cap).MatchString(arr[i]) {
			types[i] = "cap"
			numb := regexp.MustCompile(`[0-9]+`).FindString(arr[i])
			if numb == "" {
				numb = "1"
			}
			arr[i] = numb
		} else if regexp.MustCompile(up).MatchString(arr[i]) {
			types[i] = "up"
			numb := regexp.MustCompile(`[0-9]+`).FindString(arr[i])
			if numb == "" {
				numb = "1"
			}
			arr[i] = numb
		} else if regexp.MustCompile(low).MatchString(arr[i]) {
			types[i] = "low"
			numb := regexp.MustCompile(`[0-9]+`).FindString(arr[i])
			if numb == "" {
				numb = "1"
			}
			arr[i] = numb
		} else if regexp.MustCompile(hex).MatchString(arr[i]) {
			types[i] = "hex"
			numb := regexp.MustCompile(`[0-9]+`).FindString(arr[i])
			if numb == "" {
				numb = "1"
			}
			arr[i] = numb
		} else if regexp.MustCompile(word).MatchString(arr[i]) {
			types[i] = "word"
		} else if regexp.MustCompile(popen).MatchString(arr[i]) {
			types[i] = "popened"
		} else if regexp.MustCompile(pclose).MatchString(arr[i]) {
			types[i] = "pclosed"
		} else if regexp.MustCompile(pclose).MatchString(arr[i]) {
			types[i] = "pclosed"
		} else if regexp.MustCompile(sqopen).MatchString(arr[i]) {
			if sqopened {
				types[i] = "sqclosed"
				sqclosed = true
				sqopened = false
			} else if sqclosed || !(sqopened && sqclosed) {
				types[i] = "sqopened"
				sqopened = true
				sqclosed = false
			}
		} else if regexp.MustCompile(dqopen).MatchString(arr[i]) {
			if dqopened {
				types[i] = "dqclosed"
				dqclosed = true
				dqopened = false
			} else if dqclosed || !(dqopened && dqclosed) {
				types[i] = "dqopened"
				dqopened = true
				dqclosed = false
			}
		} else {
			types[i] = "other"
		}
	}
	return arr, types
}
