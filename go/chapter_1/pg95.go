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

