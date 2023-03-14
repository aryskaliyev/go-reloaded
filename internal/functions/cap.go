package functions

func Cap(word string) string {
	return Up(string(word[0])) + Low(string(word[1:]))
}
