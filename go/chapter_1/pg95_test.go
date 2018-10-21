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