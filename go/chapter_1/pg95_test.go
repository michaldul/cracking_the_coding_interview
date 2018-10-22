package pg95

import "testing"

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