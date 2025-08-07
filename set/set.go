package set

/*
Creates a set from a slice of strings
*/
func CreateSet(s []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(s))

	for _, val := range s {
		if !seen[val] {
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
}
