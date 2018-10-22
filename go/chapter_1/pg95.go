package pg95

func AllRunesUnique(s string) bool {
	var letters = make(map[rune]bool)
	for _, char := range s {
		if letters[char] == true {
			return false
		}
		letters[char] = true
	}
	return true
}

func AllRunesUniqueNoDataStructures(s string) bool {
	for i := 0; i < len(s) - 1; i++ {
		for j := i+1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}