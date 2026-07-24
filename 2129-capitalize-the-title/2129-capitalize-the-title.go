import (
	"strings"
	"unicode"
)
func capitalizeTitle(title string) string {
	words := strings.Fields(title)
	for i, word := range words {
		if len(word) <= 2 {
			words[i] = strings.ToLower(word)
			continue
		}
		runes := []rune(word)
		for j, val := range runes {
			if j == 0 {
				runes[j] = unicode.ToUpper(val)
				continue
			}
			runes[j] = unicode.ToLower(val)
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}