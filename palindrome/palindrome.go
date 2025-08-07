package palindrome

import (
	"regexp"
	"strings"

	"github.com/zanpatryk/gocourse/stack"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^[:alnum:]]+`)

func CheckWithStack(data string) bool {

	cleaned := nonAlphanumericRegex.ReplaceAllString(data, "")
	cleaned = strings.ToLower(cleaned)

	var st stack.Stack

	for _, r := range cleaned {
		st.Push(r)
	}

	st.Print()

	for _, r := range cleaned {
		if st.Pop() != r {
			return false
		}
	}

	return true
}

func CheckWithDoublePointer(data string) bool {
	cleaned := nonAlphanumericRegex.ReplaceAllString(data, "")
	cleaned = strings.ToLower(cleaned)

	runes := []rune(cleaned)
	i := 0
	j := len(runes) - 1

	for i < j {
		if runes[i] != runes[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func CheckWithReversedString(data string) bool {
	cleaned := nonAlphanumericRegex.ReplaceAllString(data, "")
	cleaned = strings.ToLower(cleaned)

	runes := []rune(cleaned)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return cleaned == string(runes)
}
