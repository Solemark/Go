package fizz_buzz

import "testing"

func Test_fizz_buzz(t *testing.T){
	fizz := 3
	buzz := 5
	max := 20
	result := fizz_buzz(fizz, buzz, max)
	answer := "1\n2\nfizz\n4\nbuzz\nfizz\n7\n8\nfizz\nbuzz\n11\nfizz\n13\n14\nfizzbuzz\n16\n17\nfizz\n19\nbuzz\n"

	if result != answer {
		t.Error("input does not match answer!")
	}
}