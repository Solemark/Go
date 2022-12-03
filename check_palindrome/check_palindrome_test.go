package check_palindrome

import "testing"

func Test_palindrome_DAD(t *testing.T){
	input := "DAD"
	answer := check_palindrome(input)
	
	if !answer {
		t.Errorf("%s is not a palindrome!", input)
	}
}
