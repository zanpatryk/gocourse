package reverse

import "strings"

func ReverseCharactersOrderRaw(s string) string {
	runes := []rune(s)
	result := make([]rune, len(runes))
	i := 0

	for start := 0; start < len(runes); {
		end := start

		for end < len(runes) && runes[end] != ' ' {
			end++
		}

		for j := end - 1; j >= start; j-- {
			result[i] = runes[j]
			i++
		}

		if end < len(runes) && runes[end] == ' ' {
			result[i] = ' '
			i++
			end++
		}
		start = end
	}

	return string(result)
}

func ReverseCharactersOrder(s string) string {

	words := strings.Split(s, " ")
	for i, w := range words {
		runes := []rune(w)
		for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
			runes[l], runes[r] = runes[r], runes[l]
		}
		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}
