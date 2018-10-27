package pg174

import (
	"fmt"
	"testing"
)

func GenerateParenthesesCombinations(n int, out chan string) {
	genParenCombinations(n, 0, "", out)
	close(out)
}

func genParenCombinations(toBeOpened, toBeClosed int, path string, out chan string) {
	if toBeOpened== 0 && toBeClosed == 0 {
		out <- path
	}
	if toBeOpened > 0 {
		genParenCombinations(toBeOpened-1, toBeClosed+1, path + "(", out)
	}
	if toBeClosed > 0 {
		genParenCombinations(toBeOpened, toBeClosed-1, path + ")", out)
	}
}

func Test(t *testing.T) {
	testCases := []struct {
		n	int
	}{
		{
			n: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(string(tC.n), func(t *testing.T) {
			c := make(chan string)
			go GenerateParenthesesCombinations(tC.n, c)

			for seq := range c {
				fmt.Println(seq)
			}
		})
	}
}