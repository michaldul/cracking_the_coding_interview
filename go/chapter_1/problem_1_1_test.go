package problem_1_1

import "testing"

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

func TestAllRunesUnique(t * testing.T) {
	if !(
		AllRunesUnique("aaa") == false &&
		AllRunesUnique("a") == true &&
		AllRunesUnique("") == true) {
			t.Errorf("Err")
	}
}

func TestAllRunesUniqueNoDataStructures(t * testing.T) {
	if !(
		AllRunesUniqueNoDataStructures("aaa") == false &&
		AllRunesUniqueNoDataStructures("aba") == false &&
		AllRunesUniqueNoDataStructures("a") == true &&
		AllRunesUniqueNoDataStructures("") == true) {
			t.Errorf("Err")
	}
}